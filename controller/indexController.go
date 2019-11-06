package controller

import (
	"gin-inject/service"

	"github.com/gin-gonic/gin"
)

//Index 注入IStartService
type Index struct {
	Service service.IStartService `inject:""`
}

//GetName 调用IStartService的Say方法
func (i *Index) GetName(ctx *gin.Context) {
	var message = ctx.Param("msg")
	Json(ctx, &Res{200, i.Service.Say(message)})
}
