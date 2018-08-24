package comm

import (
	"github.com/gin-gonic/gin"
	"gocl/library/error"
	"gocl/library/system"
)

func NotFound(ctx *gin.Context) {
	io := sys.IO{Ctx: ctx}
	io.Response(e.NOTFOUND, gin.H{
		"ip":   io.Ip(),
		"path": io.Request().URL.Path,
	})
}

