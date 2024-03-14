// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// RefundInfo is the golang structure for table refund_info.
type RefundInfo struct {
	Id        int         `json:"id"        ` // 售后退款表
	Number    string      `json:"number"    ` // 售后订单号
	OrderId   int         `json:"orderId"   ` // 订单id
	GoodsId   int         `json:"goodsId"   ` // 要售后的商品id
	Reason    string      `json:"reason"    ` // 退款原因
	Status    int         `json:"status"    ` // 状态 1待处理 2同意退款 3拒绝退款
	UserId    int         `json:"userId"    ` // 用户id
	CreatedAt *gtime.Time `json:"createdAt" ` //
	UpdatedAt *gtime.Time `json:"updatedAt" ` //
	DeletedAt *gtime.Time `json:"deletedAt" ` //
}