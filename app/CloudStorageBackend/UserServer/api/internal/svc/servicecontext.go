package svc

import (
	"UserServer/api/internal/config"
	"UserServer/api/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config
	model.CssUserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:       c,
		CssUserModel: model.NewCssUserModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
	}
}
