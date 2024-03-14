// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PositionInfo is the golang structure for table position_info.
type PositionInfo struct {
	Id        int         `json:"id"        ` //
	PicUrl    string      `json:"picUrl"    ` // 图片链接
	GoodsName string      `json:"goodsName" ` // 商品名称
	Link      string      `json:"link"      ` // 跳转链接
	Sort      int         `json:"sort"      ` // 排序
	GoodsId   int         `json:"goodsId"   ` // 商品id
	CreatedAt *gtime.Time `json:"createdAt" ` //
	UpdatedAt *gtime.Time `json:"updatedAt" ` //
	DeletedAt *gtime.Time `json:"deletedAt" ` //
}
