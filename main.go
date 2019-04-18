package main

import (
	"fmt"
	"github.com/mangmang/app"
	"github.com/mangmang/models"
	"github.com/mangmang/pkg/gredis"
	"github.com/mangmang/pkg/setting"
	"net/http"
)

func main() {
	setting.Setup()
	models.Setup()
	gredis.Setup()

	router := app.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.ServerSetting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ServerSetting.ReadTimeout,
		WriteTimeout:   setting.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	_ = s.ListenAndServe()
}
