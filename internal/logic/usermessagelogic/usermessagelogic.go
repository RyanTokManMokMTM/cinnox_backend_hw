package usermessagelogic

import (
	"github.com/gin-gonic/gin"
	"github.com/ryantokmanmokmtm/cinnox_backend_hw/global"
	"github.com/ryantokmanmokmtm/cinnox_backend_hw/internal/svc"
	"github.com/ryantokmanmokmtm/cinnox_backend_hw/internal/types"
)

type UserMessageLogicService struct {
	ctx    *gin.Context
	svcCtx *svc.ServiceContext
}

func NewUserMessageLogicService(svcCtx *svc.ServiceContext, ctx *gin.Context) *UserMessageLogicService {
	return &UserMessageLogicService{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (s *UserMessageLogicService) GetUserMessage(req *types.UserMessageReq) (*types.UserMessageResp, error) {
	global.Log.WithField("UserToken", req.UserToken).Debug("Fetching User Message")

	resp, err := s.svcCtx.DAO.FindAll(s.ctx, req.UserToken)
	if err != nil {
		return nil, err
	}

	global.Log.Debugf("message : %+v", resp)
	var messageList []*types.UserMessage
	for _, v := range *resp {
		messageList = append(messageList, &types.UserMessage{
			Message: v.UserMessage,
			Time:    v.DefaultModel.StartTime,
		})
	}

	return &types.UserMessageResp{
		UserMessages: messageList,
	}, nil
}
