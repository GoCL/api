package main

import (
	"fmt"
	"net/http"
	"gocl/router"
	"gocl/config"
)

func main() {
	routers := router.Init()
	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", app.HttpPort),
		Handler:        routers,
		ReadTimeout:    app.ReadTimeout,
		WriteTimeout:   app.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	server.ListenAndServe()
}
