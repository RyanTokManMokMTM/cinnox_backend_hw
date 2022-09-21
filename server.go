package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/ryantokmanmokmtm/cinnox_backend_hw/global"
	"github.com/ryantokmanmokmtm/cinnox_backend_hw/internal/handler"
	"github.com/ryantokmanmokmtm/cinnox_backend_hw/pkg/setting"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	"os"
)

//init server
var channelToken = "FzEby/2ql8wgNnobUFgE6L58Lk0UljxNOZGg21ff//fo6BZ3aUP2288SBEB6gMdywPy6W0LSUVWSWRqB90fojXDSl5erBMDI1Nelpn0/6h9Fii7zPT4NNwGPSsYshPVFv0p4KQ1i9W2Ofp4vzzydogdB04t89/1O/w1cDnyilFU="
var ChannelSecret = "3916ee36c100f7be534460f34cc3d3b3"

func init() {

	if err := loadSetting(); err != nil {
		log.Fatalln(err.Error())
	}
	if err := setUpLineBot(); err != nil {
		log.Fatalln(err.Error())
	}

	//Setup Logger
	global.Log = logrus.New()
	global.Log.SetLevel(logrus.DebugLevel)
	global.Log.SetFormatter(&logrus.JSONFormatter{})
	global.Log.SetOutput(os.Stdout)

}

func main() {
	global.Log.WithFields(logrus.Fields{
		"Server Name": "Demo",
	}).Info("starting the server")

	gin.SetMode(global.ServerSetting.Mode)
	route := handler.NewRoute()
	server := http.Server{
		Addr:    fmt.Sprintf("%s:%d", global.ServerSetting.Host, global.ServerSetting.Port),
		Handler: route,
	}

	log.Fatalln(server.ListenAndServe())
}

func setUpLineBot() error {
	var err error
	global.Line, err = linebot.New(ChannelSecret, channelToken)
	if err != nil {
		return err
	}
	return nil
}

func loadSetting() error {
	s, err := setting.NewSetting()
	if err != nil {
		return nil
	}

	if err := s.ReadSection("Server", &global.ServerSetting); err != nil {
		return err
	}

	return nil
}
