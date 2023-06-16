package model

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"strings"
)

var _ ConvertConfigModel = (*customConvertConfigModel)(nil)

type (
	// ConvertConfigModel is an interface to be customized, add more methods here,
	// and implement the added methods in customConvertConfigModel.
	ConvertConfigModel interface {
		convertConfigModel

		FindAll(ctx context.Context) (res []*ConvertConfig, err error)
	}

	customConvertConfigModel struct {
		*defaultConvertConfigModel
	}
)

func (c *customConvertConfigModel) FindAll(ctx context.Context) (res []*ConvertConfig, err error) {
	s, args, err := c.rowBuilder().Where("deleted_at is null").ToSql()
	if err != nil {
		return nil, err
	}

	err = c.conn.QueryRowsCtx(ctx, &res, s, args...)
	return
}

// NewConvertConfigModel returns a model for the database table.
func NewConvertConfigModel(conn sqlx.SqlConn) ConvertConfigModel {
	return &customConvertConfigModel{
		defaultConvertConfigModel: newConvertConfigModel(conn),
	}
}

// export logic
func (c *customConvertConfigModel) insertBuilder() squirrel.InsertBuilder {
	return squirrel.Insert(c.table).Columns(strings.Split(convertConfigRowsExpectAutoSet, ",")...)
}

// export logic
func (c *customConvertConfigModel) rowBuilder() squirrel.SelectBuilder {
	return squirrel.Select(convertConfigRows).From(c.table)
}

// export logic
func (c *customConvertConfigModel) countBuilder(field string) squirrel.SelectBuilder {
	return squirrel.Select("COUNT(" + field + ")").From(c.table)
}

// export logic
func (c *customConvertConfigModel) sumBuilder(field string) squirrel.SelectBuilder {
	return squirrel.Select("IFNULL(SUM(" + field + "),0)").From(c.table)
}
