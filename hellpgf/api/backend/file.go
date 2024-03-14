package backend

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type FileUploadReq struct {
	//ime:"multipart/form-data"标签用于指定结构体字段的MIME类型为multipart/form-data，
	//通常用于处理包含文件上传的HTTP表单数据。
	g.Meta `path:"/file" method:"post" mime:"multipart/form-data" tags:"工具" dc:"上传文件"`
	//参数接收的数据类型使用 *ghttp.UploadFile 如果需要接口文档也支持文件类型，那么参数的标签中设置type类型也为file
	File *ghttp.UploadFile `json:"file" type:"file" dc:"选择上传文件"`
}

type FileUploadRes struct {
	Name string `json:"name" dc:"文件名称"`
	Url  string `json:"url"dc:"图片地址"`
}
