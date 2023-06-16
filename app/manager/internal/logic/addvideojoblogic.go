package logic

import (
	"context"
	"database/sql"
	"dist-encoder/app/manager/internal/model"
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

	data := &model.ConvertJob{
		InPut:  in.InPut,
		OutPut: in.OutPut,

		Status: int64(distribute.Status_Waiting),
		Host:   sql.NullString{},
		Ip:     sql.NullString{},
	}
	res, err := l.svcCtx.ConvertJobModel.Insert(l.ctx, data)
	if err != nil {
		return nil, err
	}
	id, _ := res.LastInsertId()

	return &distribute.AddVideoJobResponse{
		JobId: id,
	}, nil
}
