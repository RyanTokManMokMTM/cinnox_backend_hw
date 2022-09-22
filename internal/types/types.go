package types

type ReplyMessageReq struct {
	ReplyUserID  string `json:"user_id" binding:"required"`
	ReplyMessage string `json:"message" binding:"required"`
}

type ReplyMessageResp struct{}

type WebhookReq struct{}

type WebhookResp struct {
	//UserToken string `json:"user_token"`
}

type UserMessageReq struct {
	UserToken string `json:"user_token" form:"user_token" binding:"required"`
}

type UserMessageResp struct {
	UserMessages []*UserMessage `json:"user_message"`
}

type UserMessage struct {
	Message string `json:"message"`
	Time    int64  `json:"sent_time"`
}
