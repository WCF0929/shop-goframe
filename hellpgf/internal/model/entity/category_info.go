// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// CategoryInfo is the golang structure for table category_info.
type CategoryInfo struct {
	Id        int         `json:"id"        ` //
	ParentId  int         `json:"parentId"  ` // 父级id
	Name      string      `json:"name"      ` //
	PicUrl    string      `json:"picUrl"    ` // icon
	Level     int         `json:"level"     ` // 等级 默认1级分类
	Sort      int         `json:"sort"      ` //
	CreatedAt *gtime.Time `json:"createdAt" ` //
	UpdatedAt *gtime.Time `json:"updatedAt" ` //
	DeletedAt *gtime.Time `json:"deletedAt" ` //
}
