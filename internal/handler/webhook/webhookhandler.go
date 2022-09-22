package webhook

import (
	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/ryantokmanmokmtm/cinnox_backend_hw/global"
	"github.com/ryantokmanmokmtm/cinnox_backend_hw/internal/logic/webhooklogic"
	"github.com/ryantokmanmokmtm/cinnox_backend_hw/internal/svc"
	"net/http"
)

func WebhookHandler(svcCtx *svc.ServiceContext) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		events, err := global.Line.ParseRequest(ctx.Request)
		if err != nil {
			if err == linebot.ErrInvalidSignature {
				global.Log.Error("LineBot ErrInvalidSignature")
				ctx.AbortWithStatus(http.StatusBadRequest)
			} else {
				global.Log.Error("LineBot err : %v ", err.Error())
				ctx.AbortWithStatus(http.StatusInternalServerError)
			}
			return
		}

		service := webhooklogic.NewWebhookLogic(svcCtx, ctx)
		_, err = service.EventHandler(events)

		if err != nil {
			global.Log.Error("LineBot Webhook err :%v ", err.Error())
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		} else {
			ctx.AbortWithStatus(http.StatusOK)
		}
	}
}
