package logic

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/ryantokmanmokmtm/cinnox_backend_hw/global"
	"github.com/ryantokmanmokmtm/cinnox_backend_hw/internal/svc"
	"net/http"
)

type WebhookLogicService struct {
	ctx    *gin.Context
	svcCtx *svc.ServiceContext
}

func NewWebhookLogic(svcCtx *svc.ServiceContext, ctx *gin.Context) *WebhookLogicService {
	return &WebhookLogicService{
		svcCtx: svcCtx,
		ctx:    ctx,
	}
}

func (s *WebhookLogicService) EventHandler(events []*linebot.Event) {
	for _, event := range events {
		switch msg := event.Message.(type) {
		case *linebot.TextMessage:
			global.Log.
				WithField("Received Message From ", msg.ID).
				Info(fmt.Printf("message : %s", msg.Text))

			//Calling Dao and sending back to user
			if err := s.svcCtx.DAO.InsertOne(s.ctx, event.Source.UserID, msg.Text, event.Timestamp.Unix()); err != nil {
				global.Log.WithField("EventHandler", "").Error(err)
				s.ctx.AbortWithStatus(http.StatusInternalServerError)
				return
			}

			resp, err := global.Line.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(fmt.Sprintf("Message recevid : %s", msg.Text))).Do()
			if err != nil {
				global.Log.WithField("Reply Message To", msg.ID).Error(err)
				s.ctx.AbortWithStatus(http.StatusInternalServerError)
				return
			}

			global.Log.
				WithField("Reply Message To ", msg.ID).
				Info(fmt.Printf("succeed and resp %+v", resp))

			s.ctx.AbortWithStatus(http.StatusOK)
		default:
			break
		}
	}
}
