package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"market-monitor/api/internal/config"
	"market-monitor/api/internal/handler"
	"market-monitor/api/internal/svc"
	"market-monitor/api/internal/task"
)

var configFile = flag.String("f", "api/etc/marketmonitor.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	//// 初始化 Twitter 监控
	//monitor, err := task.NewTwitterMonitor("api/internal/task/cookies.json", []string{"dwours", "bwenews"})
	//if err != nil {
	//	fmt.Printf("初始化 Twitter 监控失败: %v\n", err)
	//	return
	//}
	//monitor.Start()

	// 方程式新闻id 1483495485889564674
	// monitor, _ := task.NewTwitterMonitor("1483495485889564674")
	// monitor.Start()

	// 初始化通知器
	notifier, isLoginChan := task.NewNotifier(&task.NotifierConfig{ToGroupName: "星空社区-监控内测"})
	select {
	case isLogin := <-isLoginChan:
		if !isLogin {
			panic("微信登录失败")
		}
	}
	// 运行 Binance 公告监控
	task.StartMonitorProcess(notifier, "square", "https://www.binance.com/zh-CN/square/profile/panews", "星空社区播报：\n 来源: PANews \n")
	//task.StartMonitorProcess("square", "https://www.binance.com/zh-CN/square/profile/binance_announcement")
	task.StartMonitorProcess(notifier, "square", "https://www.binance.com/zh-CN/square/profile/bwenews", "星空社区播报：\n 来源: 方程式新闻 \n")

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
