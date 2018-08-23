package middleware

import (
	"github.com/gin-gonic/gin"
	"gocl/library"
	"gocl/config"
	"gocl/library/error"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		io := library.IO{Ctx: c}
		auth := io.Header().Get("secret")
		if auth != app.Secret {
			io.Response(e.UNAUTHORIZED)
			return
		}
		c.Next()
	}
}
