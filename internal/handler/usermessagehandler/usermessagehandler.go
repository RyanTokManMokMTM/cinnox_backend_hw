package usermessagehandler

import (
	"github.com/gin-gonic/gin"
	"github.com/ryantokmanmokmtm/cinnox_backend_hw/global"
	"github.com/ryantokmanmokmtm/cinnox_backend_hw/internal/logic/usermessagelogic"
	"github.com/ryantokmanmokmtm/cinnox_backend_hw/internal/svc"
	"github.com/ryantokmanmokmtm/cinnox_backend_hw/internal/types"
	"log"
	"net/http"
)

func UserMessageHandler(svcCtx *svc.ServiceContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var req types.UserMessageReq
		if err := ctx.ShouldBind(&req); err != nil {
			global.Log.Error("Binging: types.UserMessageReq error %v", err)
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
		log.Printf("%+v", req)
		service := usermessagelogic.NewUserMessageLogicService(svcCtx, ctx)
		resp, err := service.GetUserMessage(&req)
		if err != nil {
			global.Log.
				WithField("user_token", req.UserToken).
				Errorf("GetUserMessage err : %v", err)

			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})

			return
		}

		ctx.JSON(http.StatusOK, resp)
	}
}
