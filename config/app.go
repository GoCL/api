package app

import (
	"time"
	"fmt"
	"log"
	"github.com/go-ini/ini"
	"github.com/gin-gonic/gin"
)

var (
	Source  *ini.File
	RunMode string

	SiteName  string
	Secret string

	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
)

func init() {
	Source = loadFile("app")
	loadApp()
	gin.SetMode(RunMode)
}

func loadFile(name string) *ini.File {
	cfg, err := ini.Load(fmt.Sprintf("config/%s.ini", name))
	if err != nil {
		log.Fatalf("Fail:%v", err)
	}
	return cfg
}

func loadApp() {
	app, err := Source.GetSection("app")
	if err != nil {
		log.Fatalf("Fail to get section 'app': %v", err)
	}

	Secret = app.Key("SECRET").MustString("!@)*#)!@U#@*!@!)")
	SiteName = app.Key("SITE_NAME").MustString("Orz")
	RunMode = app.Key("RUN_MODE").MustString(gin.DebugMode)
	HttpPort = app.Key("HTTP_PORT").MustInt(8080)
	ReadTimeout = time.Duration(app.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(app.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second

}
