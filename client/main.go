package main

import (
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/knullhhf/hack22/logger"
	"github.com/knullhhf/hack22/pkg/client"
	"github.com/knullhhf/hack22/repo"
)

func main() {
	logger.InitLog()
	err := repo.InitDB()
	if err != nil {
		logs.Error("init db error:%v", err)
	}
	client.RunClient()
	beego.Run(":81")

}
