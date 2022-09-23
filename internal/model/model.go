package model

import (
	"context"
	"fmt"
	"github.com/ryantokmanmokmtm/cinnox_backend_hw/pkg/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DefaultModel struct {
	StartTime int64 `bson:"start_time"`
}

func NewMongodb(c config.Config) *mongo.Database {
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%d", c.Mongo.UserName, c.Mongo.Password, c.Mongo.Host, c.Mongo.Port)
	opt := options.Client().ApplyURI(uri)

	//connect to db
	client, err := mongo.Connect(context.TODO(), opt)
	if err != nil {
		panic(err.Error())
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		panic(err.Error())
	}

	database := client.Database(c.Mongo.Schema)
	return database
}
