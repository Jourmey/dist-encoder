package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ ConvertConfigModel = (*customConvertConfigModel)(nil)

type (
	// ConvertConfigModel is an interface to be customized, add more methods here,
	// and implement the added methods in customConvertConfigModel.
	ConvertConfigModel interface {
		convertConfigModel
	}

	customConvertConfigModel struct {
		*defaultConvertConfigModel
	}
)

// NewConvertConfigModel returns a model for the database table.
func NewConvertConfigModel(conn sqlx.SqlConn) ConvertConfigModel {
	return &customConvertConfigModel{
		defaultConvertConfigModel: newConvertConfigModel(conn),
	}
}
