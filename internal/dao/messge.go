package dao

import (
	"context"
	"github.com/ryantokmanmokmtm/cinnox_backend_hw/global"
	"github.com/ryantokmanmokmtm/cinnox_backend_hw/internal/model"
	"github.com/ryantokmanmokmtm/cinnox_backend_hw/internal/model/message"
)

func (d *DAO) InsertOne(ctx context.Context, userToken, msg string, timeStamp int64) error {
	global.Log.WithField("DAO", "Insert One").Info("Insert UserInfo and Message")

	m := &message.Message{
		UserToken:   userToken,
		UserMessage: msg,
		DefaultModel: model.DefaultModel{
			StartTime: timeStamp,
		},
	}
	collection := d.DB.Collection(m.CollectionName())
	return m.InsertOne(collection, ctx)
}

func (d *DAO) FindAll(ctx context.Context, userToken string) (*[]message.Message, error) {
	m := &message.Message{
		UserToken: userToken,
	}

	collection := d.DB.Collection(m.CollectionName())

	return m.FindAll(collection, ctx)
}
