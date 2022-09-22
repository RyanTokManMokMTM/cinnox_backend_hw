package message

import (
	"context"
	"errors"
	"github.com/ryantokmanmokmtm/cinnox_backend_hw/global"
	"github.com/ryantokmanmokmtm/cinnox_backend_hw/internal/model"
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
	global.Log.WithField("Message Model", "Insert One").Infof("%+v", m)
	collection := db.Collection(m.CollectionName())
	if collection == nil {
		return errors.New("collection not exists")
	}

	_, err := collection.InsertOne(ctx, m)
	if err != nil {
		return err
	}

	return nil
}

func (m *Message) FindOne(db *mongo.Database, ctx context.Context) ([]*Message, error) {
	collection := db.Collection(m.CollectionName())
	if collection == nil {
		return nil, errors.New("collection not exists")
	}

	var msg []*Message
	cursor, err := collection.Find(ctx, bson.M{
		"user_token": m.UserToken,
	}, options.Find().SetSkip(0))
	if err != nil {
		return nil, err
	}

	//Close the cursor
	defer func(ctx context.Context) {
		if err := cursor.Close(ctx); err != nil {
			global.Log.WithField("Find Mongodb records by user_token", m.UserToken).Error(err)
			return
		}
	}(ctx)

	//get all records
	if err := cursor.All(ctx, msg); err != nil {
		return nil, err
	}

	return msg, nil
}
