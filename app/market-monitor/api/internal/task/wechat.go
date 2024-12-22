package task

import (
	"fmt"
	"github.com/eatmoreapple/openwechat"
	"time"
)

type WechatBot struct {
	bot         *openwechat.Bot
	toGroupName string
}

func NewWechatBot(toGroupName string) *WechatBot {
	bot := openwechat.DefaultBot(openwechat.Desktop)
	return &WechatBot{
		bot:         bot,
		toGroupName: toGroupName,
	}
}

// 启动
func (wb *WechatBot) Start(isLoginChan chan bool) {
	go func() {
		//后续解耦，用生产消费
		// 注册消息处理函数
		//wb.bot.MessageHandler = func(msg *openwechat.Message) {
		//	if msg.IsText() && msg.Content == "ping" {
		//		msg.ReplyText("pong")
		//	}
		//}
		// 注册登陆二维码回调
		wb.bot.UUIDCallback = openwechat.PrintlnQrcodeUrl
		// 登陆
		if err := wb.bot.Login(); err != nil {
			fmt.Println(err)
			return
		}
		// 获取登陆的用户
		_, err := wb.bot.GetCurrentUser()
		if err != nil {
			fmt.Println(err)
			return
		}
		// 通知登录成功
		isLoginChan <- true
		// 阻塞协程
		wb.bot.Block()
	}()

}

func (wb *WechatBot) SendTextToGroup(content string) {
	user, _ := wb.bot.GetCurrentUser()
	groups, err := user.Groups()
	if err != nil {
		fmt.Println(err)
	}
	groupsAfterSearch := groups.SearchByNickName(1, wb.toGroupName)
	if groupsAfterSearch.Count() == 0 {
		println("找不到目的群组: ", wb.toGroupName)
		return
	}
	group := groupsAfterSearch.First()
	// 延迟
	time.Sleep(1 * time.Second)
	group.SendText(content)
}
