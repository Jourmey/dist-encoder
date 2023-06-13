package logic

import (
	"context"

	"dist-encoder/app/manager/internal/svc"
	"dist-encoder/pb/distribute"

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
	// todo: add your logic here and delete this line

	return &distribute.AddConvertCnfResponse{}, nil
}
