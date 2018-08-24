package controller

import (
	"github.com/gin-gonic/gin"
	"gocl/library/error"
	"gocl/library/system"
)

type IndexController struct {
}

func (ctl *IndexController) Index(ctx *gin.Context) {
	io := sys.IO{Ctx: ctx}
	// some code
	io.Show(1, 2, 3)
	io.Response(e.SUCCESS, gin.H{}, "Hello GoCL")
}

func (ctl *IndexController) Auth(ctx *gin.Context) {
	io := sys.IO{Ctx: ctx}
	// some code
	io.Response(e.SUCCESS, nil, "Auth Test")
}
