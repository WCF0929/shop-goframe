// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// CommentInfo is the golang structure for table comment_info.
type CommentInfo struct {
	Id        int         `json:"id"        ` //
	ParentId  int         `json:"parentId"  ` // 父级评论id
	UserId    int         `json:"userId"    ` //
	ObjectId  int         `json:"objectId"  ` //
	Type      int         `json:"type"      ` // 评论类型：1商品 2文章
	Content   string      `json:"content"   ` // 评论内容
	CreatedAt *gtime.Time `json:"createdAt" ` //
	UpdatedAt *gtime.Time `json:"updatedAt" ` //
	DeletedAt *gtime.Time `json:"deletedAt" ` //
}
