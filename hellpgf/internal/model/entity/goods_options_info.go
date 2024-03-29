// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// GoodsOptionsInfo is the golang structure for table goods_options_info.
type GoodsOptionsInfo struct {
	Id        int         `json:"id"        ` //
	GoodsId   int         `json:"goodsId"   ` // 商品id
	PicUrl    string      `json:"picUrl"    ` // 图片
	Name      string      `json:"name"      ` // 商品名称
	Price     int         `json:"price"     ` // 价格 单位分
	Stock     int         `json:"stock"     ` // 库存
	CreatedAt *gtime.Time `json:"createdAt" ` //
	UpdatedAt *gtime.Time `json:"updatedAt" ` //
	DeletedAt *gtime.Time `json:"deletedAt" ` //
}
