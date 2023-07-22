package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/shinhagunn/todo-backend/internal/router"
	"github.com/shinhagunn/todo-backend/pkg/logger"
	"github.com/shinhagunn/todo-backend/pkg/postgres"
	"github.com/shinhagunn/todo-backend/pkg/setting"
	"github.com/shinhagunn/todo-backend/pkg/util"
)

func init() {
	setting.Setup()
	logger.Setup()
	util.Setup()
}

func main() {
	db := postgres.Setup()

	router := router.InitRouter(db)
	readTimeout := time.Duration(setting.Cfg.Server.ReadTimeout) * time.Second
	writeTimeout := time.Duration(setting.Cfg.Server.WriteTimeout) * time.Second
	endPoint := fmt.Sprintf(":%d", setting.Cfg.Server.Port)

	server := &http.Server{
		Addr:         endPoint,
		Handler:      router,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
	}

	log.Printf("[info] start http server listening %s", endPoint)

	server.ListenAndServe()
}
