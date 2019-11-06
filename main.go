package main

import (
	"github.com/bingjian-zhu/gin-inject/cmd"
	"github.com/bingjian-zhu/gin-inject/common/setting"
	"fmt"
)

func main() {
	app := cmd.App()
	app.Run(fmt.Sprintf(":%s", setting.Config.Server.Port))
}
