package frontend

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gmeta"
)

// 分页处理
type GoodsGetListCommonReq struct {
	g.Meta `path:"/goods/list" method:"post" tags:"前台商品" summary:"查看商品"`
	CommonPaginationReq
}
type GoodsGetListCommonRes struct {
	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}
type GoodsDetailReq struct {
	g.Meta `path:"/goods/detail"method:"post"tags:"前台商品" summary:"商品详情页"`
	Id     uint `json:"id"`
}
type GoodsDetailRes struct {
	GoodsInfoBase
	Options   []GoodsOptionsBase `json:"options"` //规格sku
	Comments  []CommentBase      `json:"comments"`
	IsCollect bool               `json:"is_collect"`
}
type GoodsInfoBase struct {
	Id               int         `json:"id"               dc:""`       //
	PicUrl           string      `json:"picUrl"           dc:"图片"`     // 图片
	Name             string      `json:"name"             dc:"商品名称"`   // 商品名称
	Price            int         `json:"price"            dc:"价格"`     // 价格 单位分
	Level1CategoryId int         `json:"level1CategoryId" dc:"一级分类id"` // 1级分类id
	Level2CategoryId int         `json:"level2CategoryId" dc:"二级分类id"` // 2级分类id
	Level3CategoryId int         `json:"level3CategoryId" dc:"三级分类id"` // 3级分类id
	Brand            string      `json:"brand"            dc:"品牌"`     // 品牌
	Stock            int         `json:"stock"            dc:"库存"`     // 库存
	Sale             int         `json:"sale"             dc:"销售"`     // 销量
	Tags             string      `json:"tags"             dc:"标签"`     // 标签
	DetailInfo       string      `json:"detailInfo"       dc:"商品详情"`   // 商品详情
	CreatedAt        *gtime.Time `json:"created_at" `                  //
}
type GoodsOptionsBase struct {
	Id        int         `json:"id"        `  //
	GoodsId   int         `json:"goodsId"   `  // 商品id
	PicUrl    string      `json:"picUrl"    `  // 图片
	Name      string      `json:"name"      `  // 商品名称
	Price     int         `json:"price"     `  // 价格 单位分
	Stock     int         `json:"stock"     `  // 库存
	CreatedAt *gtime.Time `json:"created_at" ` //
}
type BaseGoodsColumns struct {
	gmeta.Meta `orm:"table:goods_info"`
	Id         string `json:"id"`
	Name       string `json:"name"`
	Price      int    `json:"price"`
	Brand      string `json:"brand"`
	Tags       string `json:"tags"`
	PicUrl     string `json:"pic_url"`
	DetailInfo string `json:"detail_info"`
}
