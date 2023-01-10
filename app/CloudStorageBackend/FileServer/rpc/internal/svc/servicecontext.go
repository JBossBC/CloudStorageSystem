package svc

import (
	"cloudStorageSystem/app/CloudStorageBackend/FileServer/model/PojoDB/fileMetaPojo"
	"cloudStorageSystem/app/CloudStorageBackend/FileServer/rpc/internal/config"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config    config.Config
	FileModel fileMetaPojo.FilemetatableModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.DB.DataSource)
	return &ServiceContext{
		Config:    c,
		FileModel: fileMetaPojo.NewFilemetatableModel(sqlConn, c.Cache),
	}
}
