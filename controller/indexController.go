package controller

import (
	"gin-inject/service"

	"github.com/gin-gonic/gin"
)

type Index struct {
	Service service.IStartService `inject:""`
}

func (i *Index) GetName(ctx *gin.Context) {
	var message= ctx.Param("msg")
	Json(ctx, &Res{200, i.Service.Say(message)})
}
