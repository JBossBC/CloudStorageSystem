// Code generated by goctl. DO NOT EDIT!

package fileMetaPojo

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	filemetatableFieldNames          = builder.RawFieldNames(&Filemetatable{})
	filemetatableRows                = strings.Join(filemetatableFieldNames, ",")
	filemetatableRowsExpectAutoSet   = strings.Join(stringx.Remove(filemetatableFieldNames, "`updated_at`", "`update_time`", "`create_at`", "`created_at`", "`create_time`", "`update_at`"), ",")
	filemetatableRowsWithPlaceHolder = strings.Join(stringx.Remove(filemetatableFieldNames, "`creator`", "`updated_at`", "`update_time`", "`create_at`", "`created_at`", "`create_time`", "`update_at`"), "=?,") + "=?"

	cacheCloudStorageSystemFilemetatableCreatorPrefix = "cache:cloudStorageSystem:filemetatable:creator:"
)

type (
	filemetatableModel interface {
		Insert(ctx context.Context, data *Filemetatable) (sql.Result, error)
		FindOne(ctx context.Context, creator string, name string) (*Filemetatable, error)
		Update(ctx context.Context, data *Filemetatable) error
		Delete(ctx context.Context, creator string, name string) error
		Query(ctx context.Context, creator string) ([]*Filemetatable, error)
		DeleteHard(ctx context.Context, timeInt time.Time) error
	}

	defaultFilemetatableModel struct {
		sqlc.CachedConn
		table string
	}

	Filemetatable struct {
		Creator     string       `db:"creator"`
		CreateGroup string       `db:"createGroup"`
		Name        string       `db:"name"`
		Description string       `db:"description"`
		CreateTime  time.Time    `db:"create_time"`
		Authority   string       `db:"authority"`
		TypeOf      string       `db:"typeOf"`
		UpdateTime  time.Time    `db:"update_time"`
		Size        int64        `db:"size"`
		IsDir       int          `db:"isDir"`
		DeleteTime  sql.NullTime `db:"delete_time"`
	}
)

func newFilemetatableModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultFilemetatableModel {
	return &defaultFilemetatableModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`filemetatable`",
	}
}

func (m *defaultFilemetatableModel) Delete(ctx context.Context, creator string, name string) error {
	cloudStorageSystemFilemetatableCreatorKey := fmt.Sprintf("%s%v%v", cacheCloudStorageSystemFilemetatableCreatorPrefix, creator, name)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `creator` = ? and `name`= ? ", m.table)
		return conn.ExecCtx(ctx, query, creator, name)
	}, cloudStorageSystemFilemetatableCreatorKey)
	return err
}

func (m *defaultFilemetatableModel) FindOne(ctx context.Context, creator string, name string) (*Filemetatable, error) {
	cloudStorageSystemFilemetatableCreatorKey := fmt.Sprintf("%s%v%v", cacheCloudStorageSystemFilemetatableCreatorPrefix, creator, name)
	var resp Filemetatable
	err := m.QueryRowCtx(ctx, &resp, cloudStorageSystemFilemetatableCreatorKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `creator` = ? and `name` = ? limit 1", filemetatableRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, creator, name)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
func (m *defaultFilemetatableModel) Query(ctx context.Context, creator string) ([]*Filemetatable, error) {
	var resp []*Filemetatable
	err := m.QueryRowsNoCacheCtx(ctx, &resp, fmt.Sprintf("select %s from %s where `creator` = '%s'", filemetatableRows, m.table, creator))
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
func (m *defaultFilemetatableModel) DeleteHard(ctx context.Context, timeInt time.Time) error {
	err := m.CachedConn.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		var resp []*Filemetatable
		err := session.QueryRowsCtx(ctx, &resp, fmt.Sprintf("select %s from %s where `delete_time` < \"%s\" for update", filemetatableRows, m.table, timeInt))
		if err != nil {
			if err == sqlc.ErrNotFound {
				return nil
			}
			return err
		}
		for _, value := range resp {
			_, err := session.ExecCtx(ctx, fmt.Sprintf("delete from %s where 'creator'= %s and 'name' = %s ", m.table, value.Creator, value.Name))
			if err != nil {
				fmt.Printf("delete hard the filemeta info error:%v", err)
				return err
			}
			logx.Info(fmt.Sprintf("successful delete hard the filemetinfo------> creator:%s , filename:%s", value.Creator, value.Name))
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
func (m *defaultFilemetatableModel) Insert(ctx context.Context, data *Filemetatable) (sql.Result, error) {
	cloudStorageSystemFilemetatableCreatorKey := fmt.Sprintf("%s%v:%v", cacheCloudStorageSystemFilemetatableCreatorPrefix, data.Creator, data.Name)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, filemetatableRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Creator, data.CreateGroup, data.Name, data.Description, data.Authority, data.TypeOf, data.Size, data.IsDir, data.DeleteTime)
	}, cloudStorageSystemFilemetatableCreatorKey)
	return ret, err
}

func (m *defaultFilemetatableModel) Update(ctx context.Context, data *Filemetatable) error {
	cloudStorageSystemFilemetatableCreatorKey := fmt.Sprintf("%s%v:%v", cacheCloudStorageSystemFilemetatableCreatorPrefix, data.Creator, data.Name)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `creator` = ?", m.table, filemetatableRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.CreateGroup, data.Name, data.Authority, data.TypeOf, data.Size, data.IsDir, data.DeleteTime, data.Creator)
	}, cloudStorageSystemFilemetatableCreatorKey)
	return err
}
func (m *defaultFilemetatableModel) query(ctx context.Context, creator string) ([]Filemetatable, error) {
	return nil, nil
}
func (m *defaultFilemetatableModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheCloudStorageSystemFilemetatableCreatorPrefix, primary)
}

func (m *defaultFilemetatableModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `creator` = ? limit 1", filemetatableRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultFilemetatableModel) tableName() string {
	return m.table
}
