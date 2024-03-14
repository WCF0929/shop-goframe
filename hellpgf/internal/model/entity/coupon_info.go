// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// CouponInfo is the golang structure for table coupon_info.
type CouponInfo struct {
	Id         int         `json:"id"         ` //
	Name       string      `json:"name"       ` //
	Price      int         `json:"price"      ` // 优惠前面值 单位分
	GoodsIds   string      `json:"goodsIds"   ` // 关联使用的goods_ids  逗号分隔
	CategoryId int         `json:"categoryId" ` // 关联使用的分类id
	CreatedAt  *gtime.Time `json:"createdAt"  ` //
	UpdatedAt  *gtime.Time `json:"updatedAt"  ` //
	DeletedAt  *gtime.Time `json:"deletedAt"  ` //
}
