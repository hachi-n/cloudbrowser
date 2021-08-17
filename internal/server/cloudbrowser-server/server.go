package cloudbrowser_server

import (
	"github.com/gin-gonic/gin"
	"github.com/hachi-n/cloudbrowser/internal/handlers/ec2"
	"github.com/hachi-n/cloudbrowser/internal/server/internal/initializer"
	"github.com/hachi-n/cloudbrowser/internal/server/internal/middleware"
	_ "github.com/hachi-n/cloudbrowser/pack/assets"
)

func StartDaemon() error {
	engine := gin.Default()
	engine.Use(middleware.ServerLogFormat)
	//engine.Static("/assets", "./assets")

	initializer.Initialize(engine)
	engine.GET("/ec2", ec2.Index)
	engine.Run(":3000")

	return nil
}
