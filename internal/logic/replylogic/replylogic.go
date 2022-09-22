package replylogic

import (
	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/ryantokmanmokmtm/cinnox_backend_hw/global"
	"github.com/ryantokmanmokmtm/cinnox_backend_hw/internal/svc"
	"github.com/ryantokmanmokmtm/cinnox_backend_hw/internal/types"
	"github.com/sirupsen/logrus"
)

type ReplyLogicService struct {
	ctx    *gin.Context
	svcCtx *svc.ServiceContext
}

func NewReplyLogicService(svcCtx *svc.ServiceContext, ctx *gin.Context) *ReplyLogicService {
	return &ReplyLogicService{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (s *ReplyLogicService) ReplyMessage(req *types.ReplyMessageReq) (*types.ReplyMessageResp, error) {
	//TODO: Reply to Line Bot
	resp, err := global.Line.PushMessage(req.ReplyUserID, linebot.NewTextMessage(req.ReplyMessage)).Do()
	if err != nil {
		global.Log.WithFields(
			logrus.Fields{
				"reply_token":   req.ReplyUserID,
				"reply_message": req.ReplyMessage,
			}).Error(err)
		return nil, err
	}

	global.Log.WithFields(
		logrus.Fields{
			"reply_token":   req.ReplyUserID,
			"reply_message": req.ReplyMessage,
		}).Infof("reply to user succeeded : %+v", resp)

	return &types.ReplyMessageResp{}, nil
}
