package logic

import (
	"context"

	"dist-encoder/app/manager/internal/svc"
	"dist-encoder/pb/distribute"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddVideoJobLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddVideoJobLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddVideoJobLogic {
	return &AddVideoJobLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddVideoJob 添加视频转码任务
func (l *AddVideoJobLogic) AddVideoJob(in *distribute.AddVideoJobRequest) (*distribute.AddVideoJobResponse, error) {
	// todo: add your logic here and delete this line

	return &distribute.AddVideoJobResponse{}, nil
}
