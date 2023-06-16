package logic

import (
	"context"

	"dist-encoder/app/manager/internal/svc"
	"dist-encoder/pb/distribute"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetVideoJobResultLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSetVideoJobResultLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetVideoJobResultLogic {
	return &SetVideoJobResultLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 工作机器上报工作状态
func (l *SetVideoJobResultLogic) SetVideoJobResult(in *distribute.SetVideoJobResultRequest) (*distribute.SetVideoJobResultResponse, error) {

	_, err := l.svcCtx.ConvertJobModel.UpdateStatus(l.ctx, in.JobId, int64(in.Status))
	if err != nil {
		return nil, err
	}

	return &distribute.SetVideoJobResultResponse{}, nil
}
