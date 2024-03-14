package backend

import "github.com/gogf/gf/v2/frame/g"

type GoodsOptionsReq struct {
	g.Meta `path:"/good/option/add" method:"post" tags:"商品操作" summary:"添加商品操作"`
	GoodsOptionsCommonAddUpdate
}

type GoodsOptionsCommonAddUpdate struct {
	GoodsId uint   `json:"goods_id"          v:"required#商品id必填" dc:"商品id"`      //商品id
	PicUrl  string `json:"picUrl"            dc:"图片"        `                    // 图片
	Name    string `json:"name"              dc:"商品操作分类" v:"required#商品操作名称必填" ` // 商品操作名称
	Price   uint   `json:"price"             dc:"价格"    v:"required#价格必填"`       // 价格 单位分
	Stock   uint   `json:"stock"              dc:"库存" `                          // 库存
}
type GoodsOptionsRes struct {
	Id uint `json:"id"`
}

type GoodsOptionsDeleteReq struct {
	g.Meta `path:"/good/option/delete"method:"delete"tags:"商品操作"summary:"删除商品操作"`
	Id     uint `v:"min:1#请选择要删除的商品操作"dc:"商品操作id"`
}
type GoodsOptionsDeleteRes struct {
}
type GoodsOptionsUpdateReq struct {
	g.Meta `path:"/good/option/update/{Id}"method:"post" tags:"商品操作"summary:"更新商品操作"`
	Id     uint `json:"id" v:"min:1#请选择需要修改的商品操作"dc:"商品操作Id"`
	GoodsOptionsCommonAddUpdate
}
type GoodsOptionsUpdateRes struct {
	Id uint `json:"id"`
}

// 分页处理
type GoodsOptionsGetListCommonReq struct {
	g.Meta `path:"/good/option/list" method:"get" tags:"商品操作"summary:"查看商品操作"`
	CommonPaginationReq
}
type GoodsOptionsGetListCommonRes struct {
	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}
