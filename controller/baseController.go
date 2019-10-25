package controller

import "github.com/gin-gonic/gin"

type Res struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

func Json(ctx *gin.Context, res *Res) {
	ctx.JSON(res.Code, &res)
}
