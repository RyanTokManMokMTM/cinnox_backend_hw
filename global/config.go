package global

import (
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/ryantokmanmokmtm/cinnox_backend_hw/pkg/setting"
	"github.com/sirupsen/logrus"
)

var (
	ServerSetting *setting.ServerSection
	Log           *logrus.Logger
	Line          *linebot.Client
)
