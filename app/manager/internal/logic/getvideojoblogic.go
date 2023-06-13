package logic

import (
	"context"

	"dist-encoder/app/manager/internal/svc"
	"dist-encoder/pb/distribute"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetVideoJobLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetVideoJobLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetVideoJobLogic {
	return &GetVideoJobLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 工作机器获取工作任务
func (l *GetVideoJobLogic) GetVideoJob(in *distribute.GetVideoJobRequest) (*distribute.GetVideoJobResponse, error) {
	// todo: add your logic here and delete this line

	return &distribute.GetVideoJobResponse{}, nil
}
