// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// GoodsInfo is the golang structure for table goods_info.
type GoodsInfo struct {
	Id               int         `json:"id"               ` //
	PicUrl           string      `json:"picUrl"           ` // 图片
	Name             string      `json:"name"             ` // 商品名称
	Price            int         `json:"price"            ` // 价格 单位分
	Level1CategoryId int         `json:"level1CategoryId" ` // 1级分类id
	Level2CategoryId int         `json:"level2CategoryId" ` // 2级分类id
	Level3CategoryId int         `json:"level3CategoryId" ` // 3级分类id
	Brand            string      `json:"brand"            ` // 品牌
	Stock            int         `json:"stock"            ` // 库存
	Sale             int         `json:"sale"             ` // 销量
	Tags             string      `json:"tags"             ` // 标签
	DetailInfo       string      `json:"detailInfo"       ` // 商品详情
	CreatedAt        *gtime.Time `json:"createdAt"        ` //
	UpdatedAt        *gtime.Time `json:"updatedAt"        ` //
	DeletedAt        *gtime.Time `json:"deletedAt"        ` //
}
