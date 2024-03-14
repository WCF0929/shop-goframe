package frontend

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type AddCommentReq struct {
	g.Meta   `path:"/comment/add" method:"post" tags:"前台评论"summary:"添加评论"`
	ParentId uint   `json:"parent_id"  dc:"父级id"`
	UserId   int    `json:"userId"    dc:"用户id"`                       // 用户id
	ObjectId uint   `json:"objectId"dc:"对象id" v:"required#评论的对象id必填" ` // 对象id
	Contents string `json:"contents"dc:"评论内容"`
	Type     uint8  `json:"type"   dc:"评论类型:1商品 2文章"  v:"in:1,2"` // 评论类型：1商品 2文章
}
type AddCommentRes struct {
	Id uint `json:"id"`
}

type DeleteCommentReq struct {
	g.Meta `path:"/comment/delete" method:"post" tags:"前台评论"summary:"删除评论"`
	Id     uint `json:"id"`
}
type DeleteCommentRes struct {
	Id uint `json:"id"`
}
type ListCommentReq struct {
	g.Meta `path:"comment/list" method:"post" tags:"前台评论" summary:"评论列表"`
	Type   uint8 `json:"type"v:"in:0,1,2," dc:"评论类型"`
	CommonPaginationReq
}

type ListCommentRes struct {
	Page  int         `json:"page"description:"分页码"`
	Size  int         `json:"size"description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
	List  interface{} `json:"list" description:"列表"`
}
type ListCommentItem struct {
	Id       uint        `json:"id"`
	UserId   int         `json:"user_id"    ` // 用户id
	ObjectId int         `json:"object_id"  ` // 对象id
	Type     int         `json:"type"      `  // 评论类型：1.商品 2.文章
	Goods    interface{} `json:"goods" `
	Article  interface{} `json:"article"`
}

type CommentBase struct {
	Id        int          `json:"id"        `  //
	ParentId  int          `json:"parent_id"  ` // 父级评论id
	UserId    int          `json:"user_id"    ` //
	User      UserInfoBase `json:"user"dc:"发布评论的用户信息"`
	ObjectId  int          `json:"object_id"  ` //
	Type      int          `json:"type"      `  // 评论类型：1商品 2文章
	Content   string       `json:"content"   `  // 评论内容
	CreatedAt *gtime.Time  `json:"created_at" ` //

}
