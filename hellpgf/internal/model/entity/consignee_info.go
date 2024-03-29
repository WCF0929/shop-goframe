// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ConsigneeInfo is the golang structure for table consignee_info.
type ConsigneeInfo struct {
	Id        int         `json:"id"        ` // 收货地址表
	UserId    int         `json:"userId"    ` //
	IsDefault int         `json:"isDefault" ` // 默认地址1  非默认0
	Name      string      `json:"name"      ` //
	Phone     string      `json:"phone"     ` //
	Province  string      `json:"province"  ` //
	City      string      `json:"city"      ` //
	Town      string      `json:"town"      ` // 县区
	Street    string      `json:"street"    ` // 街道乡镇
	Detail    string      `json:"detail"    ` // 地址详情
	CreatedAt *gtime.Time `json:"createdAt" ` //
	UpdatedAt *gtime.Time `json:"updatedAt" ` //
	DeletedAt *gtime.Time `json:"deletedAt" ` //
}
