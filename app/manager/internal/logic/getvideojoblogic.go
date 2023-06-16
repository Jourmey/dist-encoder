package logic

import (
	"context"
	"dist-encoder/app/manager/internal/model"
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

	job, err := l.svcCtx.ConvertJobModel.FindOneByStatus(l.ctx, int64(distribute.Status_Waiting))
	if err == model.ErrNotFound {
		return &distribute.GetVideoJobResponse{}, nil
	} else if err != nil {
		return nil, err
	}

	_, err = l.svcCtx.ConvertJobModel.UpdateStatusAndHost(l.ctx, job.Id, int64(distribute.Status_Doing), in.Host, in.Ip)
	if err != nil {
		return nil, err
	}

	resp := &distribute.GetVideoJobResponse{
		Job: &distribute.VideoJob{
			JobId:     job.Id,
			InPut:     job.InPut,
			OutPut:    job.OutPut,
			ConvertId: job.ConvertId,
			Status:    distribute.Status(job.Status),
		},
		ConvertCnf: nil,
	}

	if job.ConvertId != 0 {

		conf, err := l.svcCtx.ConvertConfigModel.FindOne(l.ctx, job.ConvertId)
		if err != nil {
			return nil, err
		}

		resp.ConvertCnf = transformConvertCnf(conf)
	}

	return resp, nil
}
