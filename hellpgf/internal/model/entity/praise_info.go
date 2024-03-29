// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PraiseInfo is the golang structure for table praise_info.
type PraiseInfo struct {
	Id        int         `json:"id"        ` // 点赞表
	UserId    int         `json:"userId"    ` //
	Type      int         `json:"type"      ` // 点赞类型 1商品 2文章
	ObjectId  int         `json:"objectId"  ` // 点赞对象id 方便后期扩展
	CreatedAt *gtime.Time `json:"createdAt" ` //
	UpdatedAt *gtime.Time `json:"updatedAt" ` //
}
