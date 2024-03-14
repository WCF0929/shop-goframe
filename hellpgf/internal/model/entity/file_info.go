// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// FileInfo is the golang structure for table file_info.
type FileInfo struct {
	Id        int         `json:"id"        ` //
	Name      string      `json:"name"      ` // 图片名称
	Src       string      `json:"src"       ` // 本地文件存储路径
	Url       string      `json:"url"       ` // URL地址
	UserId    int         `json:"userId"    ` // 用户id
	CreatedAt *gtime.Time `json:"createdAt" ` //
	UpdatedAt *gtime.Time `json:"updatedAt" ` //
}
