package common

import (
	"github.com/gin-gonic/gin"
	"gocl/library"
	"gocl/library/error"
)

func NotFound(ctx *gin.Context) {
	io := library.IO{Ctx: ctx}
	io.Response(e.NOTFOUND, gin.H{"path": io.Request().URL.Path})
}
