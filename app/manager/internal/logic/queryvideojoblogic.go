package logic

import (
	"context"

	"dist-encoder/app/manager/internal/svc"
	"dist-encoder/pb/distribute"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryVideoJobLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryVideoJobLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryVideoJobLogic {
	return &QueryVideoJobLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryVideoJob 查询转码任务
func (l *QueryVideoJobLogic) QueryVideoJob(in *distribute.QueryVideoJobRequest) (*distribute.QueryVideoJobResponse, error) {
	// todo: add your logic here and delete this line

	return &distribute.QueryVideoJobResponse{}, nil
}
