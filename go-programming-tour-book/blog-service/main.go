package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"net/http"
	"time"
	"yazidchen.com/go-programming-tour-book/blog-service/global"
	"yazidchen.com/go-programming-tour-book/blog-service/internal/base"
	"yazidchen.com/go-programming-tour-book/blog-service/internal/routers"
	"yazidchen.com/go-programming-tour-book/blog-service/pkg/setting"
)

func init() {
	err := setupSetting()
	if err != nil {
		log.Fatal().Err(err).Msg("init.setupSetting err")
	}
	err = setupDBEngine()
	if err != nil {
		log.Fatal().Err(err).Msg("init.setupDBEngine err")
	}
}

func main() {
	gin.SetMode(global.ServerSetting.RunMode)
	r := routers.NewRouter()
	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        r,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20, // 1048576 Byte = 1024 KB = 1 M
	}
	_ = s.ListenAndServe()
}

func setupSetting() error {
	s, err := setting.NewSetting()
	if err != nil {
		return err
	}
	err = s.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}

	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	return nil
}

func setupDBEngine() error {
	var err error
	global.DBEngine, err = base.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}

	return nil
}
