package main

import (
	"flag"
	"fmt"
	"net/http"
	"user/api/internal/config"
	"user/api/internal/handler"
	"user/api/internal/svc"
	"user/entity"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "api/etc/user.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	// 跨域配置
	domains := []string{"*", "http://127.0.0.1", "http://localhost"}
	server := rest.MustNewServer(
		c.RestConf,
		rest.WithCors(domains...),
		rest.WithCustomCors(func(header http.Header) {
			header.Add("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token,Authorization,Token,X-Token,X-User-Id,OS,Platform, Version")
			header.Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS,PATCH")
			header.Set("Access-Control-Expose-Headers", "Content-Length, Content-Type")
		}, nil, "*"),
	)

	//server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	// 自动映射实体到数据库
	ctx.DB.AutoMigrate(&entity.Users{}, &entity.Assets{})

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
