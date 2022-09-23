package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/ryantokmanmokmtm/cinnox_backend_hw/global"
	"github.com/ryantokmanmokmtm/cinnox_backend_hw/internal/handler"
	"github.com/ryantokmanmokmtm/cinnox_backend_hw/internal/svc"
	"github.com/ryantokmanmokmtm/cinnox_backend_hw/pkg/config"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	"os"
)

//init server
const channelToken = "FzEby/2ql8wgNnobUFgE6L58Lk0UljxNOZGg21ff//fo6BZ3aUP2288SBEB6gMdywPy6W0LSUVWSWRqB90fojXDSl5erBMDI1Nelpn0/6h9Fii7zPT4NNwGPSsYshPVFv0p4KQ1i9W2Ofp4vzzydogdB04t89/1O/w1cDnyilFU="
const channelSecret = "3916ee36c100f7be534460f34cc3d3b3"

var configFileName = flag.String("f", "config_local.yaml", "the config file")

func init() {
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
	flag.Parse()
	log.Println(*configFileName)

	var c config.Config
	if err := loadSetting(&c); err != nil {
		panic(err.Error())
	}

	global.Log.WithFields(logrus.Fields{
		"Server Name": "Demo",
	}).Info("starting the server")

	//Server Setting
	gin.SetMode(c.Server.Mode)
	ctx := svc.NewServiceContext(c)
	route := handler.RegisterHandlers(ctx)
	server := http.Server{
		Addr:    fmt.Sprintf("%s:%d", c.Server.Host, c.Server.Port),
		Handler: route,
	}

	log.Fatalln(server.ListenAndServe())
}

//SetUp Function
func setUpLineBot() error {
	var err error
	global.Line, err = linebot.New(channelSecret, channelToken)
	if err != nil {
		return err
	}
	return nil
}

func loadSetting(c *config.Config) error {
	//load file

	s, err := config.NewSetting(*configFileName)
	if err != nil {
		return nil
	}
	var serverSetting *config.ServerSection
	var mongoSetting *config.DBSection

	if err := s.ReadSection("Server", &serverSetting); err != nil {
		return err
	}

	if err := s.ReadSection("DB", &mongoSetting); err != nil {
		return err
	}

	c.Server = serverSetting
	c.Mongo = mongoSetting

	return nil
}
