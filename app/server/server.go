package server

import (
	orderHnadler "github.com/o-ga09/tutorial-go-fr/app/server/handler/order"
	"gofr.dev/pkg/gofr"
)

func Run() {
	app := gofr.New()
	handler := orderHnadler.OrderHandler{}
    app.GET("/v1/health",gofr.HeartBeatHandler)
    app.GET("/v1/health_db",gofr.HealthHandler)
	app.GET("/v1/order",handler.OrderHandler)
    app.Start()
}