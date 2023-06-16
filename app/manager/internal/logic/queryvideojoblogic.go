package logic

import (
	"context"
	"errors"

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

	if in.Page == nil {
		return nil, errors.New("invalid page")
	}

	data, err := l.svcCtx.ConvertJobModel.Query(l.ctx, in.Page.OrderBy, in.Page.Offset, in.Page.Limit)
	if err != nil {
		return nil, err
	}

	jobs := make([]*distribute.VideoJob, 0, len(data))
	for i := 0; i < len(data); i++ {
		jobs = append(jobs, &distribute.VideoJob{
			JobId:     data[i].Id,
			InPut:     data[i].InPut,
			OutPut:    data[i].OutPut,
			ConvertId: data[i].ConvertId,
			Status:    distribute.Status(data[i].Status),
		})
	}

	return &distribute.QueryVideoJobResponse{
		Page: in.Page,
		Jobs: jobs,
	}, nil
}
