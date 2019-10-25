package main

import (
	"gin-inject/cmd"
	"gin-inject/common/setting"
	"fmt"
)

func main() {
	app := cmd.App()
	app.Run(fmt.Sprintf(":%s", setting.Config.Server.Port))
}
