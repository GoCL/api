package controller

import (
	"github.com/gin-gonic/gin"
	"gocl/library"
	"gocl/library/error"
)

type IndexController struct {
}

func (ctl *IndexController) Index(ctx *gin.Context) {
	io := library.IO{Ctx: ctx}
	// some code
	io.Show(1, 2, 3)
	io.Response(e.SUCCESS, gin.H{}, "Hello GoCL")
}

func (ctl *IndexController) Auth(ctx *gin.Context) {
	io := library.IO{Ctx: ctx}
	// some code
	io.Response(e.SUCCESS, gin.H{}, "Auth Test")
}
