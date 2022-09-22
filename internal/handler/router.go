package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/ryantokmanmokmtm/cinnox_backend_hw/internal/handler/webhook"
	"github.com/ryantokmanmokmtm/cinnox_backend_hw/internal/svc"
)

func RegisterHandlers(svc *svc.ServiceContext) *gin.Engine {
	engine := gin.New()
	engine.Use(gin.Recovery())
	engine.Use(gin.Logger())

	//engine.GET("/ping", func(ctx *gin.Context) {
	//	ctx.JSON(http.StatusOK, gin.H{
	//		"msg": "pong",
	//	})
	//})
	api := engine.Group("/api/v1")
	api.POST("/webhook", webhook.WebhookHandler(svc))

	return engine
}
