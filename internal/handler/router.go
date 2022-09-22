package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/ryantokmanmokmtm/cinnox_backend_hw/internal/handler/replyhandler"
	"github.com/ryantokmanmokmtm/cinnox_backend_hw/internal/handler/usermessagehandler"
	"github.com/ryantokmanmokmtm/cinnox_backend_hw/internal/handler/webhook"
	"github.com/ryantokmanmokmtm/cinnox_backend_hw/internal/svc"
	"net/http"
)

func RegisterHandlers(svc *svc.ServiceContext) *gin.Engine {
	engine := gin.New()
	engine.Use(gin.Recovery())
	engine.Use(gin.Logger())

	engine.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "pong",
		})
		return
	})

	api := engine.Group("/api/v1")
	api.POST("/webhook", webhook.WebhookHandler(svc))
	api.POST("/message/reply", replyhandler.ReplayHandler(svc))
	api.GET("/messages/get", usermessagehandler.UserMessageHandler(svc))

	return engine
}
