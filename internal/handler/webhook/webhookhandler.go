package webhook

import (
	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/ryantokmanmokmtm/cinnox_backend_hw/global"
	"github.com/sirupsen/logrus"
	"net/http"
)

func WebhookHandler() func(ctx *gin.Context) {
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

		for _, event := range events {
			//check event type
			if event.Type == linebot.EventTypeMessage {
				switch msg := event.Message.(type) {
				case *linebot.TextMessage:
					global.Log.WithFields(logrus.Fields{
						"Line Event": "text message",
					}).Info("received message %s", msg.Text)

					if _, err := global.Line.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("Hello Testing")).Do(); err != nil {
						global.Log.WithFields(logrus.Fields{
							"Line Event": "text message",
						}).Info(err)

						ctx.AbortWithStatus(http.StatusInternalServerError)
					}
				}
			}
		}
	}
}
