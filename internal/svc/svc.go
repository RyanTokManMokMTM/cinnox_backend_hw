package svc

import (
	"github.com/ryantokmanmokmtm/cinnox_backend_hw/internal/dao"
	"github.com/ryantokmanmokmtm/cinnox_backend_hw/internal/model"
	"github.com/ryantokmanmokmtm/cinnox_backend_hw/pkg/config"
)

type ServiceContext struct {
	//Put all Config here
	Config config.Config
	DAO    *dao.DAO //USING DATABASE ACCESS AS MODEL LAYER
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		DAO:    dao.NewDao(model.NewMongodb(c)),
	}
}
