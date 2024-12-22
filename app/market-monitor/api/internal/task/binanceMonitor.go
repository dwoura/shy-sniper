package task

import (
	"context"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
	"log"
	"sort"
	"strings"
	"time"
)

type BinanceMonitor struct {
	TargetUrl                  string   // 轮询目标网址
	DataIDs                    []string // 获取到的 ID 组
	NewestDataID               string   // 最新消息
	BinanceAnnouncementMsgChan chan string
	BrowserCtx                 context.Context
	CancelFunc                 context.CancelFunc
}

func NewBinanceMonitor(mode string, targetUrl string) *BinanceMonitor {
	ctx, cancel := chromedp.NewContext(context.Background())
	if mode == "square" {
		return &BinanceMonitor{
			//targetUrl:                  "https://www.binance.com/zh-CN/square/profile/binance_announcement",
			//targetUrl: "https://www.binance.com/zh-CN/square/profile/panews",
			//targetUrl:                  "https://www.binance.com/zh-CN/square/profile/square-creator-46016c811",
			TargetUrl:                  targetUrl,
			DataIDs:                    []string{},
			NewestDataID:               "",
			BinanceAnnouncementMsgChan: make(chan string, 5),
			BrowserCtx:                 ctx,
			CancelFunc:                 cancel,
		}
	}
	panic("Unsupported mode: " + mode)
}

func (bm *BinanceMonitor) Start() {
	go bm.poll()
}

// 轮询
func (bm *BinanceMonitor) poll() {
	for range time.Tick(10 * time.Second) {
		var statusCodes []int
		// 定义 chromedp 任务
		var dataIDs []string
		var html string
		// 执行 chromedp 任务
		err := chromedp.Run(bm.BrowserCtx,
			network.Enable(),
			network.SetCacheDisabled(true),
			chromedp.Navigate(bm.TargetUrl),
			chromedp.ActionFunc(func(ctx context.Context) error {
				// 监听网络请求
				chromedp.ListenTarget(ctx, func(ev interface{}) {
					if event, ok := ev.(*network.EventResponseReceived); ok {
						statusCodes = append(statusCodes, int(event.Response.Status))
					}
				})
				return nil
			}),
			chromedp.Evaluate(`Array.from(document.querySelectorAll('div[class="FeedList css-vurnku"] > div[class="FeedItem css-1y7889g"]')).map(item => item.getAttribute('data-id'))`, &dataIDs), // 获取 data-id
			chromedp.Evaluate(`window.scrollBy(0, window.innerHeight);`, nil), // 模拟滚动
			chromedp.Sleep(2*time.Second),                                     // 等待页面加载完成
			chromedp.OuterHTML(`div[class="FeedList css-vurnku"]`, &html),     // 获取目标元素的 HTML
		)

		if err != nil {
			log.Printf("访问失败: %v", err)
			log.Printf("等待30s重试")
			time.Sleep(30 * time.Second)
			continue
		}

		for _, code := range statusCodes {
			if code == 403 || code == 429 {
				log.Printf("检测到错误码 %d，更新上下文...", code)
				bm.updateContext() // 在更新中进入另一个函数
				return             // 关闭此函数
			}
		}

		// 加载 HTML
		doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
		if err != nil {
			log.Println("加载 HTML 出错", err)
		}

		// 解析 HTML 获取所有子元素的 data-id
		err = bm.updateLatestDataID(doc)
		if err != nil {
			log.Println("Error parsing data-id:", err)
		}

		// 有新公告
		if bm.DataIDs[len(bm.DataIDs)-1] != bm.NewestDataID {

			selection := doc.Find("a[href='/square/post/" + bm.DataIDs[len(bm.DataIDs)-1] + "']")
			if selection.Length() > 0 {
				bm.NewestDataID = bm.DataIDs[len(bm.DataIDs)-1]
				selection.Each(func(i int, s *goquery.Selection) {
					title := s.Text()
					link, _ := s.Attr("href")
					//log.Printf("新公告：%s - %s\n", title, "https://www.binance.com"+link)
					result := time.Now().String() + ": " + title + "\n" + "https://www.binance.com" + link
					// 生产数据
					bm.BinanceAnnouncementMsgChan <- result
				})
			} else {
				log.Println(time.Now().String(), "未能从该文章 id 找到新公告")
			}

		} else {
			fmt.Println(time.Now().String(), "无新公告")
		}

	}
}

// 更新Chromedp上下文
func (bm *BinanceMonitor) updateContext() {
	bm.CancelFunc()
	ctx, cancel := chromedp.NewContext(context.Background())
	bm.BrowserCtx = ctx
	bm.CancelFunc = cancel
	go bm.poll()
}

// 解析 HTML，找到最新的 data-id
func (bm *BinanceMonitor) updateLatestDataID(doc *goquery.Document) error {
	// 查找目标元素
	var tmpDataIDs []string
	// 获取到这个元素集，获取该元素集的前五个 data-id，判断是否有新 data-id
	doc.Find(`div[class="FeedBuzzBaseView_FeedBuzzBaseViewRoot__1sC8Q FeedBuzzBaseViewRoot ltr"]`).Each(func(i int, s *goquery.Selection) {
		id, exists := s.Attr("data-id")
		if !exists {
			println("No data-id found in element")
			return
		}
		tmpDataIDs = append(tmpDataIDs, id)
	})
	sort.Strings(tmpDataIDs)

	// 更新 top5DataIDs
	bm.DataIDs = tmpDataIDs
	return nil
}
