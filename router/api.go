package router

import (
	"github.com/gin-gonic/gin"
	"gocl/middleware"
	"gocl/controller"
	"gocl/library/common"
)

/**
 * 控制器声明
 */
var (
	idx controller.IndexController
)

func Init() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())   // 开起日志中间件
	router.Use(gin.Recovery()) // 开起异常恢复的中间件

	/**
	 * 普通路由
	 */
	router.GET("/", idx.Index)

	/**
	 * 以下为鉴权中间件组路由
	 */
	authRouter := router.Group("")
	authRouter.Use(middleware.Auth())
	{
		auth(authRouter)
	}

	/**
	 * Not Found
	 */
	router.NoRoute(common.NotFound)
	return router
}
