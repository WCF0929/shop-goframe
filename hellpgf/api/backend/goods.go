package backend

import "github.com/gogf/gf/v2/frame/g"

type GoodsReq struct {
	g.Meta `path:"/goods/add" method:"post" tags:"商品" summary:"添加商品"`
	GoodsCommonAddUpdate
}

type GoodsCommonAddUpdate struct {
	PicUrl           string `json:"picUrl"            dc:"图片"        `                // 图片
	Name             string `json:"name"              dc:"商品分类" v:"required#商品名称必填" ` // 商品名称
	Price            int    `json:"price"             dc:"价格"    v:"required#价格必填"`   // 价格 单位分
	Level1CategoryId int    `json:"level1CategoryId"  dc:"1级分类id"`                    // 1级分类id
	Level2CategoryId int    `json:"level2CategoryId"  dc:"2级分类id"`                    // 2级分类id
	Level3CategoryId int    `json:"level3CategoryId"  dc:"3级分类id"`                    // 3级分类id
	Brand            string `json:"brand"             dc:"品牌"`                        // 品牌
	Stock            int    `json:"stock"             dc:"库存" `                       // 库存
	Sale             int    `json:"sale"              dc:"销量" `                       // 销量
	Tags             string `json:"tags"              dc:"标签" `                       // 标签
	DetailInfo       string `json:"detailInfo"        dc:"商品详情" `
}
type GoodsRes struct {
	Id uint `json:"id"`
}

type GoodsDeleteReq struct {
	g.Meta `path:"/goods/delete"method:"delete"tags:"商品"summary:"删除商品"`
	Id     uint `v:"min:1#请选择要删除的商品"dc:"商品id"`
}
type GoodsDeleteRes struct {
}
type GoodsUpdateReq struct {
	g.Meta `path:"/goods/update/{Id}"method:"post" tags:"商品"summary:"更新商品"`
	Id     uint `json:"id" v:"min:1#请选择需要修改的商品"dc:"商品Id"`
	GoodsCommonAddUpdate
}
type GoodsUpdateRes struct {
	Id uint `json:"id"`
}

// 分页处理
type GoodsGetListCommonReq struct {
	g.Meta `path:"/goods/list" method:"get" tags:"商品"summary:"查看商品"`
	CommonPaginationReq
}
type GoodsGetListCommonRes struct {
	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}
