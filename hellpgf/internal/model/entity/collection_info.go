// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// CollectionInfo is the golang structure for table collection_info.
type CollectionInfo struct {
	Id        int         `json:"id"        ` //
	UserId    int         `json:"userId"    ` // 用户id
	ObjectId  int         `json:"objectId"  ` // 对象id
	Type      int         `json:"type"      ` // 收藏类型：1商品 2文章
	CreatedAt *gtime.Time `json:"createdAt" ` //
	UpdatedAt *gtime.Time `json:"updatedAt" ` //
}
