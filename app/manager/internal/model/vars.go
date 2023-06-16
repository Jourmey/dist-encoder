package model

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var ErrNotFound = sqlx.ErrNotFound

func OpenSqlite() (conn sqlx.SqlConn, err error) {

	conn = sqlx.NewSqlConn("sqlite3", "etc/manager.db")

	return
}
