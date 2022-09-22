package webhooklogic

import (
	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/ryantokmanmokmtm/cinnox_backend_hw/global"
	"github.com/ryantokmanmokmtm/cinnox_backend_hw/internal/svc"
	"github.com/ryantokmanmokmtm/cinnox_backend_hw/internal/types"
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

func (s *WebhookLogicService) EventHandler(events []*linebot.Event) (*types.WebhookResp, error) {
	//TODO: Send back the reply token
	for _, event := range events {
		switch msg := event.Message.(type) {
		case *linebot.TextMessage:
			//Calling Dao and sending back to user
			if err := s.svcCtx.DAO.InsertOne(s.ctx, event.Source.UserID, msg.Text, event.Timestamp.Unix()); err != nil {
				global.Log.WithField("EventHandler", "").Error(err)
				return nil, err
			}

			return &types.WebhookResp{}, nil

		default:
			break
		}
	}

	return &types.WebhookResp{}, nil
}
