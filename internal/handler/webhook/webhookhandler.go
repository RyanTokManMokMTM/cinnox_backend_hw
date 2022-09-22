package webhook

import (
	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/ryantokmanmokmtm/cinnox_backend_hw/global"
	"github.com/ryantokmanmokmtm/cinnox_backend_hw/internal/logic"
	"github.com/ryantokmanmokmtm/cinnox_backend_hw/internal/svc"
	"net/http"
)

func WebhookHandler(svcCtx *svc.ServiceContext) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		events, err := global.Line.ParseRequest(ctx.Request)
		if err != nil {
			if err == linebot.ErrInvalidSignature {
				ctx.AbortWithStatus(http.StatusBadRequest)
			} else {
				ctx.AbortWithStatus(http.StatusInternalServerError)
			}
			return
		}

		service := logic.NewWebhookLogic(svcCtx, ctx)
		service.EventHandler(events)

	}
}
