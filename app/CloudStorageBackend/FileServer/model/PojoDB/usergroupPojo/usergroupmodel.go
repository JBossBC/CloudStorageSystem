package usergroupPojo

import (
	"github.com/lib/pq"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UsergroupModel = (*customUsergroupModel)(nil)

type (
	// UsergroupModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUsergroupModel.
	UsergroupModel interface {
		usergroupModel
	}

	customUsergroupModel struct {
		*defaultUsergroupModel
	}
)

// NewUsergroupModel returns a model for the database table.
func NewUsergroupModel(conn sqlx.SqlConn, c cache.CacheConf) UsergroupModel {
	return &customUsergroupModel{
		defaultUsergroupModel: newUsergroupModel(conn, c),
	}
}
