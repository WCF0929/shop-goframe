package model

import (
	"github.com/gogf/gf/v2/os/gtime"
	"hellpgf/internal/model/do"
)

type AddCommentInput struct {
	ParentId uint   `json:"parent_id"  dc:"父级id"`
	UserId   uint   `json:"userId"    dc:"用户id"`                       // 用户id
	ObjectId uint   `json:"objectId"dc:"对象id" v:"required#评论的对象id必填" ` // 对象id
	Type     uint8  `json:"type"   dc:"评论类型:1商品 2文章"  v:"in:1,2"`      //
	Content  string `json:"content"dc:"内容"`
}

type AddCommentOutput struct {
	Id uint `json:"id"`
}
type DeleteCommentInput struct {
	Id uint `json:"id"`
}

// DeleteCommentOutput 删除点赞返回
type DeleteCommentOutput struct {
	Id uint `json:"id"`
}

// CommentListInput 获取内容列表
type CommentListInput struct {
	Page int   // 分页号码
	Size int   // 分页数量，最大50
	Type uint8 // 排序类型(0:最新, 默认。1:活跃, 2:热度)
}

// CommentListOutputItem 查询列表分类
type CommentListOutputItem struct {
	Id        uint        `json:"id"`
	UserId    int         `json:"user_id"    ` // 用户id
	ObjectId  int         `json:"object_id"  ` // 对象id
	Type      int         `json:"type"      `  // 评论类型：1.商品 2.文章
	Content   string      `json:"content"dc:"内容"`
	ParentId  uint        `json:"parent_id"  dc:"父级id"`
	Goods     GoodsItem   `json:"goods" orm:"with:id=object_id"`
	Article   ArticleItem `json:"article"orm:"with:id=object_id "`
	CreatedAt *gtime.Time `json:"created_at"` // 创建时间
	UpdatedAt *gtime.Time `json:"updated_at"` // 修改时间
}

// CommentListOutput 查询列表结果
type CommentListOutput struct {
	List  []CommentListOutputItem `json:"list" description:"列表"`
	Page  int                     `json:"page" description:"分页码"`
	Size  int                     `json:"size" description:"分页数量"`
	Total int                     `json:"total" description:"数据总数"`
}

// CheckIsCommentInput 校验当前用户是否评论
type CheckIsCommentInput struct {
	UserId   uint
	ObjectId uint
	Type     uint8
}

type CommentBase struct {
	do.CommentInfo
	User UserInfoBase `json:"user" orm:"with:id=user_id"`
}
