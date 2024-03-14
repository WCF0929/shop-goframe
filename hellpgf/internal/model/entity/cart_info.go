// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// CartInfo is the golang structure for table cart_info.
type CartInfo struct {
	Id             int         `json:"id"             ` // 购物车表
	UserId         int         `json:"userId"         ` //
	GoodsOptionsId int         `json:"goodsOptionsId" ` // 商品规格id
	Count          int         `json:"count"          ` // 商品数量
	CreatedAt      *gtime.Time `json:"createdAt"      ` //
	UpdatedAt      *gtime.Time `json:"updatedAt"      ` //
	DeletedAt      *gtime.Time `json:"deletedAt"      ` //
}
