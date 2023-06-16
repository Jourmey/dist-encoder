package logic

import (
	"context"

	"dist-encoder/app/manager/internal/svc"
	"dist-encoder/pb/distribute"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListConvertCnfLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListConvertCnfLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListConvertCnfLogic {
	return &ListConvertCnfLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// ListConvertCnf 查询转码任务
func (l *ListConvertCnfLogic) ListConvertCnf(in *distribute.ListConvertCnfRequest) (*distribute.ListConvertCnfResponse, error) {

	res, err := l.svcCtx.ConvertConfigModel.FindAll(l.ctx)
	if err != nil {
		return nil, err
	}

	converts := make([]*distribute.ConvertCnf, 0, len(res))

	for i := 0; i < len(res); i++ {

		converts = append(converts, transformConvertCnf(res[i]))
	}

	return &distribute.ListConvertCnfResponse{
		Converts: converts,
	}, nil
}
