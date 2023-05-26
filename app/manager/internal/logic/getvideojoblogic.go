package logic

import (
	"context"

	"dist-encoder/app/manager/internal/svc"
	"dist-encoder/pb/manager"

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

func (l *GetVideoJobLogic) GetVideoJob(in *manager.Worker) (*manager.VideoJob, error) {
	// todo: add your logic here and delete this line

	return &manager.VideoJob{}, nil
}
