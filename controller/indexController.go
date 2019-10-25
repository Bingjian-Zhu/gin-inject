package controller

import (
	"gin-inject/service"

	"github.com/gin-gonic/gin"
)

type Index struct {
	Service service.IStartService `inject:""`
}

func (i *Index) GetName(ctx *gin.Context) {
	var message string
	ctx.ShouldBind(&message)
	Json(ctx, &Res{200, i.Service.Say(message)})
}

func (i *Index)GetID(ctx *gin.Context)  {
	var ID = 10
	Json(ctx, &Res{200, i.Service.GetID(ID)})
}
