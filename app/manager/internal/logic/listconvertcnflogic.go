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
	// todo: add your logic here and delete this line

	return &distribute.ListConvertCnfResponse{}, nil
}
