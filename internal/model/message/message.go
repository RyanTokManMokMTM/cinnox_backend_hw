package message

import (
	"context"
	"errors"
	"github.com/ryantokmanmokmtm/cinnox_backend_hw/global"
	"github.com/ryantokmanmokmtm/cinnox_backend_hw/internal/model"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Message struct {
	UserToken   string `bson:"user_token"`
	UserMessage string `bson:"message"`
	model.DefaultModel
}

func (m *Message) CollectionName() string {
	return "message"
}

func (m *Message) InsertOne(db *mongo.Database, ctx context.Context) error {
	collection := db.Collection(m.CollectionName())
	if collection == nil {
		global.Log.Errorf("Collection %s not exists.", m.CollectionName())
		return errors.New("collection not exists")
	}

	_, err := collection.InsertOne(ctx, m)
	if err != nil {
		global.Log.WithFields(logrus.Fields{
			"user_token":   m.UserToken,
			"user_message": m.UserMessage,
			"create_time":  m.DefaultModel.StartTime,
		}).Errorf("Insert Failed : %v", err)
		return err
	}

	return nil
}

func (m *Message) FindAll(db *mongo.Database, ctx context.Context) (*[]Message, error) {
	collection := db.Collection(m.CollectionName())
	if collection == nil {
		return nil, errors.New("collection not exists")
	}

	msg := new([]Message)
	cursor, err := collection.Find(ctx, bson.M{
		"user_token": m.UserToken,
	}, options.Find().SetSkip(0))

	if err != nil {
		global.Log.WithFields(logrus.Fields{
			"user_token": m.UserToken,
		}).Errorf("Find User Err : %v", err)
		return nil, err
	}

	//Close the cursor
	defer func(ctx context.Context) {
		if err := cursor.Close(ctx); err != nil {
			global.Log.Errorf("FindAll - Closed cursor err : %v", err)
			return
		}
	}(ctx)

	//get all records
	if err := cursor.All(ctx, msg); err != nil {
		global.Log.Errorf("FindAll - Get All Records err : %v", err)
		return nil, err
	}

	return msg, nil
}
