package backend

import "github.com/gogf/gf/v2/frame/g"

type CategoryReq struct {
	g.Meta `path:"/category/add" tags:"商品分类" method:"post" summary:"商品分类"`
	CommonAddUpdate
}
type CommonAddUpdate struct {
	ParentId uint   `json:"parent_id" v:"required#parent_id不能为空" dc:"父级id"`
	Name     string `json:"name"    v:"required#name不能为空" dc:"商品名称"`
	PicUrl   string `json:"pic_url"    v:"required#商品链接不能为空" dc:"图片链接"` //冗余设计
	Level    uint8  `json:"level"     dc:"默认分类 一级分类"`
	Sort     uint8  `json:"sort"    dc:"排序"`
}
type CategoryRes struct {
	CategoryId uint `json:"category_id"`
}

type CategoryDeleteReq struct {
	g.Meta   `path:"/category/delete"   method:"delete"  tags:"商品分类"summary:"删除内容接口"`
	Id       uint   `v:"min:1#请选择要删除的商品分类" dc:"商品分类图id"`
	ParentId uint   `json:"parent_id" v:"required#parent_id不能为空" dc:"父级id"`
	Name     string `json:"name"    v:"required#name不能为空" dc:"商品名称"`
	PicUrl   string `json:"pic_url"    v:"required#商品链接不能为空" dc:"图片链接"` //冗余设计
	Level    uint8  `json:"level"     dc:"默认分类 一级分类"`
	Sort     uint8  `json:"sort"    dc:"排序"`
}
type CategoryDeleteRes struct {
}

type CategoryUpdateReq struct {
	g.Meta `path:"/category/update/{Id}" method:"post" tags:"商品分类" summary:"修改内容接口"`
	Id     uint `json:"id"      v:"min:1#请选择需要修改的内容" dc:"内容Id"`
	CommonAddUpdate
}
type CategoryUpdateRes struct {
	Id uint `json:"id"`
}
type CategoryGetListCommonReq struct {
	g.Meta `path:"/category/list" method:"get" tags:"商品分类" summary:"商品分类列表接口"`
	Sort   int `json:"sort"   in:"query" dc:"排序类型"`
	CommonPaginationReq
}
type CategoryGetListCommonRes struct {
	//g.Meta `mime:"text/html" type:"string" example:"<html/>"`
	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}
type CategoryGetListAllCommonReq struct {
	g.Meta `path:"/category/list/all" method:"get" tags:"商品分类" summary:"商品分类全部分类"`
}
type CategoryGetListAllCommonRes struct {
	//g.Meta `mime:"text/html" type:"string" example:"<html/>"`
	List  interface{} `json:"list" description:"列表"`
	Total int         `json:"total" description:"数据总数"`
}
