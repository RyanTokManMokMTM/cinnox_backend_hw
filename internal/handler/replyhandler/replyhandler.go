package replyhandler

import (
	"github.com/gin-gonic/gin"
	"github.com/ryantokmanmokmtm/cinnox_backend_hw/global"
	"github.com/ryantokmanmokmtm/cinnox_backend_hw/internal/logic/replylogic"
	"github.com/ryantokmanmokmtm/cinnox_backend_hw/internal/svc"
	"github.com/ryantokmanmokmtm/cinnox_backend_hw/internal/types"
	"github.com/sirupsen/logrus"
	"net/http"
)

func ReplayHandler(svcCtx *svc.ServiceContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.ReplyMessageReq
		if err := ctx.ShouldBind(&req); err != nil {
			global.Log.Error("Binding types.ReplyMessageReq error %v", err.Error())
			return
		}

		service := replylogic.NewReplyLogicService(svcCtx, ctx)
		resp, err := service.ReplyMessage(&req)
		if err != nil {
			global.Log.
				WithFields(logrus.Fields{
					"user_id": req.ReplyUserID,
					"message": req.ReplyMessage,
				}).Error("ReplyMessage err : %v", err)

			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, resp)
	}
}
