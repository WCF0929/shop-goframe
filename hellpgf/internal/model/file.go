package model

import "github.com/gogf/gf/v2/net/ghttp"

type FileUploadInput struct {
	//上传文件对象
	File       *ghttp.UploadFile
	Name       string
	RandomName bool
}

type FileUploadOutput struct {
	Id   uint //返回id
	Name string
	Src  string
	Url  string
}
