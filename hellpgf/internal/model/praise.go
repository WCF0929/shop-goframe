package model

import (
	"github.com/gogf/gf/v2/os/gtime"
)

type AddPraiseInput struct {
	UserId   uint  `json:"userId"    dc:"用户id"`                       // 用户id
	ObjectId uint  `json:"objectId"dc:"对象id" v:"required#收藏的对象id必填" ` // 对象id
	Type     uint8 `json:"type"   dc:"收藏类型:1商品 2文章"  v:"in:1,2"`      //
}

type AddPraiseOutput struct {
	Id uint `json:"id"`
}
type DeletePraiseInput struct {
	Id       uint  `json:"id"`
	UserId   uint  `json:"userId"    dc:"用户id"`                       // 用户id
	ObjectId uint  `json:"objectId"dc:"对象id" v:"required#收藏的对象id必填" ` // 对象id
	Type     uint8 `json:"type"   dc:"收藏类型:1商品 2文章"  v:"in:1,2"`      //
}

// DeletePraiseOutput 删除点赞返回
type DeletePraiseOutput struct {
	Id uint `json:"id"`
}

// PraiseListInput 获取内容列表
type PraiseListInput struct {
	Page int   // 分页号码
	Size int   // 分页数量，最大50
	Type uint8 // 排序类型(0:最新, 默认。1:活跃, 2:热度)
}

// PraiseListOutputItem 查询列表分类
type PraiseListOutputItem struct {
	Id        uint        `json:"id"`
	UserId    int         `json:"user_id"    ` // 用户id
	ObjectId  int         `json:"object_id"  ` // 对象id
	Type      int         `json:"type"      `  // 收藏类型：1.商品 2.文章
	Goods     GoodsItem   `json:"goods" orm:"with:id=object_id"`
	Article   ArticleItem `json:"article"orm:"with:id=object_id "`
	CreatedAt *gtime.Time `json:"created_at"` // 创建时间
	UpdatedAt *gtime.Time `json:"updated_at"` // 修改时间
}

// PraiseListOutput 查询列表结果
type PraiseListOutput struct {
	List  []PraiseListOutputItem `json:"list" description:"列表"`
	Page  int                    `json:"page" description:"分页码"`
	Size  int                    `json:"size" description:"分页数量"`
	Total int                    `json:"total" description:"数据总数"`
}

// CheckIsPraiseInput 校验当前用户是否收藏
type CheckIsPraiseInput struct {
	UserId   uint
	ObjectId uint
	Type     uint8
}
