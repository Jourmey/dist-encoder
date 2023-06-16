package logic

import (
	"dist-encoder/app/manager/internal/model"
	"dist-encoder/pb/distribute"
	"github.com/zeromicro/go-zero/core/jsonx"
)

func transformConvertCnf(conf *model.ConvertConfig) *distribute.ConvertCnf {
	if conf == nil {
		return nil
	}
	var (
		inKwArgs  []*distribute.KwArgs
		outKwArgs []*distribute.KwArgs
	)

	if conf.InArgs.Valid {
		_ = jsonx.UnmarshalFromString(conf.InArgs.String, &inKwArgs)
	}

	if conf.OutArgs.Valid {
		_ = jsonx.UnmarshalFromString(conf.OutArgs.String, &outKwArgs)
	}

	resp := &distribute.ConvertCnf{
		CnfId:     conf.Id,
		InKwArgs:  inKwArgs,
		OutKwArgs: outKwArgs,
	}

	return resp
}
