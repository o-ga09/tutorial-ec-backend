package server

import (
	"github.com/gin-gonic/gin"
	"github.com/o-ga09/tutorial-ec-backend/app/config"
	"github.com/o-ga09/tutorial-ec-backend/app/server/handler/health"
	"github.com/o-ga09/tutorial-ec-backend/app/server/middleware"
)

func NewServer() (*gin.Engine, error) {
	r := gin.New()
	cfg, _ := config.New()
	if cfg.Env == "PROD" {
		gin.SetMode(gin.ReleaseMode)
	}

	// setting logger
	logger := middleware.New()
	httpLogger := middleware.RequestLogger(logger)

	//setting a CORS
	cors := middleware.CORS()

	r.Use(cors)
	r.Use(httpLogger)

	v1 := r.Group("/v1")
	{
		systemHandler := health.NewHealthCheckHandler()
		v1.GET("/health", systemHandler.HealthCheck)
	}

	return r, nil
}