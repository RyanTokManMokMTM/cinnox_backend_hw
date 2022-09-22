package global

import (
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/sirupsen/logrus"
)

var (
	Log  *logrus.Logger
	Line *linebot.Client
)
