package fileMetaPojo

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ FilemetatableModel = (*customFilemetatableModel)(nil)

type (
	// FilemetatableModel is an interface to be customized, add more methods here,
	// and implement the added methods in customFilemetatableModel.
	FilemetatableModel interface {
		filemetatableModel
	}

	customFilemetatableModel struct {
		*defaultFilemetatableModel
	}
)

// NewFilemetatableModel returns a model for the database table.
func NewFilemetatableModel(conn sqlx.SqlConn, c cache.CacheConf) FilemetatableModel {
	return &customFilemetatableModel{
		defaultFilemetatableModel: newFilemetatableModel(conn, c),
	}
}
