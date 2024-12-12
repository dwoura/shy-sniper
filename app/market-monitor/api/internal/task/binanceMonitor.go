package task

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
	"github.com/thoas/go-funk"
)

var targetUrl string = "https://www.binance.com/zh-CN/square/profile/binance_announcement"
var top5DataIDs []string

func runBinanceMonitor() {
	// 初始化 chromedp 上下文
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// 定义轮询间隔
	ticker := time.NewTicker(8 * time.Second)
	defer ticker.Stop()

	// 上一次抓取的最新 data-id
	var lastDataID string

	for range ticker.C {
		var html string

		// 执行 chromedp 任务
		err := chromedp.Run(ctx,
			chromedp.Navigate(targetUrl),
			chromedp.Sleep(2*time.Second),                                 // 等待页面加载完成
			chromedp.OuterHTML(`div[class="FeedList css-vurnku"]`, &html), // 获取目标元素的 HTML
		)
		if err != nil {
			log.Println("Chromedp error:", err)
			continue
		}

		// 解析 HTML 获取所有子元素的 data-id
		updatedNums := updateLatestDataID(html)
		if err != nil {
			log.Println("Error parsing data-id:", err)
			continue
		}

		for i := len(top5DataIDs) - updatedNums - 1; i < len(top5DataIDs); i++ {
			fmt.Printf("New announcement top5DataID %d: %s\n", i, top5DataIDs[i])
		}
	}
}

// 解析 HTML，找到最新的 data-id
func updateLatestDataID(html string) int {
	// 加载 HTML
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return 0
	}

	// 查找目标元素
	var latestTop5DataIDs []string
	counterNew := 0
	// 获取到这个元素集，获取该元素集的前五个 data-id，判断是否有新 data-id
	doc.Find(`div[class="FeedBuzzBaseView_FeedBuzzBaseViewRoot__1sC8Q FeedBuzzBaseViewRoot ltr"]`).Each(func(i int, s *goquery.Selection) {
		id, exists := s.Attr("data-id")
		if !exists {
			println("No data-id found in element")
			return
		}
		if funk.Contains(top5DataIDs, id) == false {
			latestTop5DataIDs = append(latestTop5DataIDs, id) // 新的接在后面，头元素移除队列
			if len(latestTop5DataIDs) > 5 {
				latestTop5DataIDs = latestTop5DataIDs[1:]
			}
			counterNew++
		}

	})
	// 更新 top5DataIDs
	top5DataIDs = latestTop5DataIDs
	return counterNew
}
