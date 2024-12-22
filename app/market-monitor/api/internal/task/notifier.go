package task

type Notifier struct {
	WechatBot *WechatBot
	// tg bot
}

type NotifierConfig struct {
	ToGroupName string
	// tg config
}

func NewNotifier(notifierConfig *NotifierConfig) (*Notifier, chan bool) {
	isLoginChan := make(chan bool)
	wechatBot := NewWechatBot(notifierConfig.ToGroupName)
	wechatBot.Start(isLoginChan)
	return &Notifier{
		WechatBot: wechatBot,
	}, isLoginChan
}

func (n *Notifier) Notify(content string) {
	if n.WechatBot != nil {
		user, _ := n.WechatBot.bot.GetCurrentUser()
		if user != nil {
			n.WechatBot.SendTextToGroup(content)
		}
	}
	// if n.TgBot != nil {}
}
