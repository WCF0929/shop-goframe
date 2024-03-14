package backend

import "github.com/gogf/gf/v2/frame/g"

type ArticleReq struct {
	g.Meta `path:"/article/add" method:"post" tags:"文章" summary:"添加文章"`
	ArticleCommonAddUpdate
}

type ArticleCommonAddUpdate struct {
	//	UserId  uint   `json:"user_id"           dc:"用户id"   v:"required#用户id必填"`
	Title   string `json:"title"             dc:"文章标题"  v:"required#名称必填"` // 图片
	PicUrl  string `json:"pic_url"           dc:"图片链接"`
	Desc    string `json:"desc"              dc:"文章摘要"   `  // 文章名称
	Praise  int    `json:"praise"            dc:"点赞"      ` // 价格 单位分
	IsAdmin uint   `d:"1"  dc:"是否是管理员"`                     // 1是默认值
	Detail  string `json:"detail"            dc:"文章详情"  v:"required#文章详情必填"`
}
type ArticleRes struct {
	Id uint `json:"id"`
}

type ArticleDeleteReq struct {
	g.Meta `path:"/article/delete"method:"delete"tags:"文章"summary:"删除文章"`
	Id     uint `v:"min:1#请选择要删除的文章"dc:"文章id"`
}
type ArticleDeleteRes struct {
}
type ArticleUpdateReq struct {
	g.Meta `path:"/article/update/{Id}"method:"post" tags:"文章"summary:"更新文章"`
	Id     uint `json:"id" v:"min:1#请选择需要修改的文章"dc:"文章Id"`
	ArticleCommonAddUpdate
}
type ArticleUpdateRes struct {
	Id uint `json:"id"`
}

// 分页处理
type ArticleGetListCommonReq struct {
	g.Meta `path:"/article/list" method:"get" tags:"文章"summary:"查看文章"`
	CommonPaginationReq
}
type ArticleGetListCommonRes struct {
	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}
