package dao

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type DAO struct {
	DB *mongo.Database
}

func NewDao(dbClient *mongo.Database) *DAO {
	return &DAO{
		DB: dbClient,
	}
}
