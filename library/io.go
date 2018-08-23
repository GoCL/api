package library

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"log"
	"time"
	"gocl/library/error"
)

type IO struct {
	Ctx *gin.Context
}

/**
 * args -> code:int
 *		   data:map[string]interface{}
 *		   msg:string
 */
func (io *IO) Response(args ...interface{}) {
	argsLen := len(args)

	var code = e.UNKNOWN
	var data = gin.H{}
	var msg = ""

	switch {
	case argsLen >= 3:
		if tmp, ok := args[2].(string); ok {
			msg = tmp
		}
		fallthrough
	case argsLen >= 2:
		if tmp, ok := args[1].(gin.H); ok {
			data = tmp
		}
		fallthrough
	case argsLen >= 1:
		if tmp, ok := args[0].(int); ok {
			code = tmp
		}
	}

	if msg == "" {
		msg = e.GetMsg(code)
	}

	io.Ctx.JSON(http.StatusOK, gin.H{
		"code": code,
		"data": data,
		"msg":  msg,
	})
	io.Ctx.Abort()
}

/**
 * args -> param:string
 *		   def:string
 */
func (io *IO) Get(args ...string) string {
	if len(args) > 1 {
		return io.Ctx.DefaultQuery(args[0], args[1])
	}
	return io.Ctx.Query(args[0])
}

/**
 * args -> param:string
 *		   def:string
 */
func (io *IO) Post(args ...string) string {
	if len(args) > 1 {
		return io.Ctx.DefaultPostForm(args[0], args[1])
	}
	return io.Ctx.PostForm(args[0])
}

func (io *IO) Header(args ...string) http.Header {
	return io.Ctx.Request.Header
}

func (io *IO) Request() *http.Request {
	return io.Ctx.Request
}

func (io *IO) Show(args ...interface{}) {
	log.Println("GoCL Debug By", time.Now().Format("2006-01-02 15:04:05"),
		"\n------------------------------")
	for _, v := range args {
		log.Println(v)
		log.Println("------------------------------")
	}
}
