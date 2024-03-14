package model

import "github.com/gogf/gf/v2/os/gtime"

// UserCouponCreateUpdateBase 创建/修改内容基类
type UserCouponCreateUpdateBase struct {
	UserId   uint  `json:"user_id"`
	CouponId uint  `json:"coupon_id"`
	Status   uint8 `json:"status"` //状态：1可用 2已用 3过期
}

// UserCouponCreateInput  创建内容
type UserCouponCreateInput struct {
	UserCouponCreateUpdateBase
	//UserId uint
}

// UserCouponCreateOutput 创建内容返回结果
type UserCouponCreateOutput struct {
	Id uint `json:"id"`
}

// UserCouponUpdateInput 修改内容
type UserCouponUpdateInput struct {
	UserCouponCreateUpdateBase
	Id uint
}

// UserCouponGetListInput 获取内容列表
type UserCouponGetListInput struct {
	Page int // 分页号码
	Size int // 分页数量，最大50
}

// UserCouponGetListOutput 查询列表结果
type UserCouponGetListOutput struct {
	List  []UserCouponGetListOutputItem `json:"list" description:"列表"`
	Page  int                           `json:"page" description:"分页码"`
	Size  int                           `json:"size" description:"分页数量"`
	Total int                           `json:"total" description:"数据总数"`
}

// UserCouponSearchInput 搜索列表
type UserCouponSearchInput struct {
	Key          string // 关键字
	Type         string // 内容模型
	UserCouponId uint   // 栏目ID
	Page         int    // 分页号码
	Size         int    // 分页数量，最大50
}

// UserCouponSearchOutput 搜索列表结果
type UserCouponSearchOutput struct {
	List  []UserCouponSearchOutputItem `json:"list"`  // 列表
	Stats map[string]int               `json:"stats"` // 搜索统计
	Page  uint                         `json:"page"`  // 分页码
	Size  uint                         `json:"size"`  // 分页数量
	Total uint                         `json:"total"` // 数据总数
}

type UserCouponGetListOutputItem struct {
	Id        uint        `json:"id"` // 自增ID
	UserId    uint        `json:"user_id"`
	CouponId  uint        `json:"coupon_id"`
	Status    uint8       `json:"status"`     //状态：1可用 2已用 3过期
	CreatedAt *gtime.Time `json:"created_at"` // 创建时间
	UpdatedAt *gtime.Time `json:"updated_at"` // 修改时间
}

type UserCouponSearchOutputItem struct {
	UserCouponGetListOutputItem
}
