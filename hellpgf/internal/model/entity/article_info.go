// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ArticleInfo is the golang structure for table article_info.
type ArticleInfo struct {
	Id        int         `json:"id"        ` //
	UserId    int         `json:"userId"    ` // 作者id
	Title     string      `json:"title"     ` // 标题
	Desc      string      `json:"desc"      ` // 摘要
	PicUrl    string      `json:"picUrl"    ` // 封面图
	IsAdmin   int         `json:"isAdmin"   ` // 1后台管理员发布 2前台用户发布
	Praise    int         `json:"praise"    ` // 点赞数
	Detail    string      `json:"detail"    ` // 文章详情
	CreatedAt *gtime.Time `json:"createdAt" ` //
	UpdatedAt *gtime.Time `json:"updatedAt" ` //
	DeletedAt *gtime.Time `json:"deletedAt" ` //
}
