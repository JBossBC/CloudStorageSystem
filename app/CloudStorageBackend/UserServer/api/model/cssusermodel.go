package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ CssUserModel = (*customCssUserModel)(nil)

type (
	// CssUserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCssUserModel.
	CssUserModel interface {
		cssUserModel
	}

	customCssUserModel struct {
		*defaultCssUserModel
	}
)

// NewCssUserModel returns a model for the database table.
func NewCssUserModel(conn sqlx.SqlConn, c cache.CacheConf) CssUserModel {
	return &customCssUserModel{
		defaultCssUserModel: newCssUserModel(conn, c),
	}
}
