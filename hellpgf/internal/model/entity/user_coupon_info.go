// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UserCouponInfo is the golang structure for table user_coupon_info.
type UserCouponInfo struct {
	Id        int         `json:"id"        ` // 用户优惠券表
	UserId    int         `json:"userId"    ` //
	CouponId  int         `json:"couponId"  ` //
	Status    int         `json:"status"    ` // 状态：1可用 2已用 3过期
	CreatedAt *gtime.Time `json:"createdAt" ` //
	UpdatedAt *gtime.Time `json:"updatedAt" ` //
	DeletedAt *gtime.Time `json:"deletedAt" ` //
}
