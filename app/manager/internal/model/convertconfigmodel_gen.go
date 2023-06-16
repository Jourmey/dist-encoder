// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	convertConfigFieldNames          = builder.RawFieldNames(&ConvertConfig{})
	convertConfigRows                = strings.Join(convertConfigFieldNames, ",")
	convertConfigRowsExpectAutoSet   = strings.Join(stringx.Remove(convertConfigFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	convertConfigRowsWithPlaceHolder = strings.Join(stringx.Remove(convertConfigFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	convertConfigModel interface {
		Insert(ctx context.Context, data *ConvertConfig) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*ConvertConfig, error)
		Update(ctx context.Context, data *ConvertConfig) error
		Delete(ctx context.Context, id int64) error
	}

	defaultConvertConfigModel struct {
		conn  sqlx.SqlConn
		table string
	}

	ConvertConfig struct {
		Id        int64          `db:"id"`         // 主健
		InArgs    sql.NullString `db:"in_args"`    // 输入kv参数配置
		OutArgs   sql.NullString `db:"out_args"`   // 输出kv参数配置
		CreatedAt time.Time      `db:"created_at"` // 创建时间
		UpdatedAt time.Time      `db:"updated_at"` // 更新时间
		DeletedAt sql.NullTime   `db:"deleted_at"`
	}
)

func newConvertConfigModel(conn sqlx.SqlConn) *defaultConvertConfigModel {
	return &defaultConvertConfigModel{
		conn:  conn,
		table: "`convert_config`",
	}
}

func (m *defaultConvertConfigModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultConvertConfigModel) FindOne(ctx context.Context, id int64) (*ConvertConfig, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", convertConfigRows, m.table)
	var resp ConvertConfig
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultConvertConfigModel) Insert(ctx context.Context, data *ConvertConfig) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, convertConfigRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.InArgs, data.OutArgs, data.DeletedAt)
	return ret, err
}

func (m *defaultConvertConfigModel) Update(ctx context.Context, data *ConvertConfig) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, convertConfigRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.InArgs, data.OutArgs, data.DeletedAt, data.Id)
	return err
}

func (m *defaultConvertConfigModel) tableName() string {
	return m.table
}
