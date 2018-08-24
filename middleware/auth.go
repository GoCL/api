package middleware

import (
	"github.com/gin-gonic/gin"
	"gocl/config"
	"gocl/library/error"
	"gocl/library/system"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		io := sys.IO{Ctx: c}
		auth := io.Header().Get("secret")
		if auth != app.Secret {
			io.Response(e.UNAUTHORIZED)
			return
		}
		c.Next()
	}
}
