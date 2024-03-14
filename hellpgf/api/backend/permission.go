package backend

import "github.com/gogf/gf/v2/frame/g"

type PermissionReq struct {
	g.Meta `path:"/permission/add"method:"post"desc:"添加权限"tags:"permission"`

	PermissionCreateUpdateBase
}

type PermissionRes struct {
	PermissionId uint `json:"permission_id"`
}
type PermissionUpdateReq struct {
	g.Meta `path:"/permission/update"method:"post"desc:"添加权限"tags:"permission"`
	Id     uint `json:"id" v:"required#id必填"`
	PermissionCreateUpdateBase
}

type PermissionCreateUpdateBase struct {
	Name string `json:"name"v:"required#名称必填" dc:"权限名称"`
	Path string `json:"desc" dc:"权限路径"`
}

type PermissionUpdateRes struct {
	Id uint `json:"id"`
}
type PermissionDeleteReq struct {
	g.Meta `path:"/permission/delete"   method:"delete"  tags:"管理员"summary:"删除内容接口"`
	Id     uint `v:"min:1#请选择要删除的权限" dc:"删除id"`
}

type PermissionDeleteRes struct {
}
type PermissionGetListCommonRes struct {
	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}
type PermissionGetListCommonReq struct {
	g.Meta `path:"/permission/list" method:"get" tags:"管理员" summary:"管理员列表接口"`
	CommonPaginationReq
}
