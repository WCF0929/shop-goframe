package controller

import (
	"context"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"hellpgf/api/backend"
	"hellpgf/internal/model"
	"hellpgf/internal/service"
)

type cFile struct {
}

var File = cFile{}

func (c *cFile) Upload(ctx context.Context, req *backend.FileUploadReq) (res *backend.FileUploadRes, err error) {
	if req.File == nil {
		//gerror.NewCode自定义错误码和错误信息
		return nil, gerror.NewCode(gcode.CodeMissingParameter, "请上传文件")
	}
	upload, err := service.File().Upload(ctx, model.FileUploadInput{
		File:       req.File,
		RandomName: true,
	})
	if err != nil {
		return nil, err
	}
	return &backend.FileUploadRes{
		Name: upload.Name,
		Url:  upload.Url,
	}, nil
}
