package model

import (
	"context"
	"database/sql"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"strings"
)

var _ ConvertJobModel = (*customConvertJobModel)(nil)

type (
	// ConvertJobModel is an interface to be customized, add more methods here,
	// and implement the added methods in customConvertJobModel.
	ConvertJobModel interface {
		convertJobModel

		FindOneByStatus(ctx context.Context, status int64) (res *ConvertJob, err error)                                      // 根据类型和状态查询单个任务
		UpdateStatus(ctx context.Context, id int64, status int64) (res sql.Result, err error)                                // 更新任务状态
		UpdateStatusAndHost(ctx context.Context, id int64, status int64, host string, ip string) (res sql.Result, err error) // 更新任务状态
		Query(ctx context.Context, orderBy string, offset int64, limit int64) (resp []*ConvertJob, err error)
	}

	customConvertJobModel struct {
		*defaultConvertJobModel
	}
)

func (c *customConvertJobModel) Query(ctx context.Context, orderBy string, offset int64, limit int64) (resp []*ConvertJob, err error) {
	s, args, err := c.rowBuilder().
		Offset(uint64(offset)).
		Limit(uint64(limit)).
		OrderBy(orderBy).ToSql()
	if err != nil {
		return nil, err
	}

	err = c.conn.QueryRowsCtx(ctx, &resp, s, args...)
	return
}

func (c *customConvertJobModel) FindOneByStatus(ctx context.Context, status int64) (res *ConvertJob, err error) {
	s, args, err := c.rowBuilder().Where("status = ?", status).Limit(1).ToSql()
	if err != nil {
		return nil, err
	}

	var resp ConvertJob
	err = c.conn.QueryRowCtx(ctx, &resp, s, args...)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (c *customConvertJobModel) UpdateStatus(ctx context.Context, id int64, status int64) (res sql.Result, err error) {
	s, args, err := squirrel.Update(c.table).
		Set("status", status).
		Where("id = ?", id).ToSql()

	return c.conn.ExecCtx(ctx, s, args...)
}

func (c *customConvertJobModel) UpdateStatusAndHost(ctx context.Context, id int64, status int64, host string, ip string) (res sql.Result, err error) {

	s, args, err := squirrel.Update(c.table).
		Set("status", status).
		Set("host", host).
		Set("ip", ip).
		Where("id = ?", id).ToSql()

	return c.conn.ExecCtx(ctx, s, args...)
}

// NewConvertJobModel returns a model for the database table.
func NewConvertJobModel(conn sqlx.SqlConn) ConvertJobModel {
	return &customConvertJobModel{
		defaultConvertJobModel: newConvertJobModel(conn),
	}
}

// export logic
func (c *customConvertJobModel) insertBuilder() squirrel.InsertBuilder {
	return squirrel.Insert(c.table).Columns(strings.Split(convertJobRowsExpectAutoSet, ",")...)
}

// export logic
func (c *customConvertJobModel) rowBuilder() squirrel.SelectBuilder {
	return squirrel.Select(convertJobRows).From(c.table)
}

// export logic
func (c *customConvertJobModel) countBuilder(field string) squirrel.SelectBuilder {
	return squirrel.Select("COUNT(" + field + ")").From(c.table)
}

// export logic
func (c *customConvertJobModel) sumBuilder(field string) squirrel.SelectBuilder {
	return squirrel.Select("IFNULL(SUM(" + field + "),0)").From(c.table)
}
