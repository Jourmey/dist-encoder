package logic

import (
	"context"
	"database/sql"
	"dist-encoder/app/manager/internal/model"
	"dist-encoder/app/manager/internal/svc"
	"dist-encoder/pb/distribute"
	"github.com/zeromicro/go-zero/core/jsonx"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddConvertCnfLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddConvertCnfLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddConvertCnfLogic {
	return &AddConvertCnfLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddConvertCnf 添加视频转码任务
func (l *AddConvertCnfLogic) AddConvertCnf(in *distribute.AddConvertCnfRequest) (*distribute.AddConvertCnfResponse, error) {
	data := &model.ConvertConfig{}

	if in.InKwArgs != nil {

		inArgs, _ := jsonx.MarshalToString(in.InKwArgs)
		data.InArgs = sql.NullString{
			String: inArgs,
			Valid:  true,
		}
	}
	if in.OutKwArgs != nil {

		outArgs, _ := jsonx.MarshalToString(in.OutKwArgs)
		data.OutArgs = sql.NullString{
			String: outArgs,
			Valid:  true,
		}
	}

	res, err := l.svcCtx.ConvertConfigModel.Insert(l.ctx, data)
	if err != nil {
		return nil, err
	}

	id, _ := res.LastInsertId()
	return &distribute.AddConvertCnfResponse{
		CnfId: id,
	}, nil
}
