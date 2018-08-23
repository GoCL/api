package router

import (
	"github.com/gin-gonic/gin"
)

/**
 * 中间件 Auth 相关路由组
 */
func auth(router *gin.RouterGroup) {

	router.GET("/auth", idx.Auth)

}
