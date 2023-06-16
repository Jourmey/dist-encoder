package svc

import (
	"dist-encoder/app/manager/internal/config"
	"dist-encoder/app/manager/internal/model"
)

type ServiceContext struct {
	Config config.Config

	ConvertJobModel    model.ConvertJobModel    // 任务表
	ConvertConfigModel model.ConvertConfigModel // 配置表
}

func NewServiceContext(c config.Config) *ServiceContext {

	conn, err := model.OpenSqlite()
	if err != nil {
		panic(err)
	}

	return &ServiceContext{
		Config: c,

		ConvertJobModel:    model.NewConvertJobModel(conn),
		ConvertConfigModel: model.NewConvertConfigModel(conn),
	}
}
