package backend

import "github.com/gogf/gf/v2/frame/g"

type UserCouponReq struct {
	g.Meta `path:"/user/coupon/add" method:"post" tags:"用户优惠券" summary:"添加用户优惠券"`
	UserCouponCommonAddUpdate
}

type UserCouponCommonAddUpdate struct {
	UserId   uint  `json:"user_id" v:"required#用户优惠券名称不能为空" dc:"用户优惠券名称"`
	CouponId uint  `json:"coupon_id" v:"required#优惠券id不能为空"dc:"可用的的商品分类"`
	Status   uint8 `json:"status"dc:"状态"`
}

type UserCouponRes struct {
	Id uint `json:"id"`
}

type UserCouponDeleteReq struct {
	g.Meta `path:"/user/coupon/delete"method:"delete"tags:"用户优惠券"summary:"删除用户优惠券"`
	Id     uint `v:"min:1#请选择要删除的用户优惠券"dc:"用户优惠券id"`
}
type UserCouponDeleteRes struct {
}
type UserCouponUpdateReq struct {
	g.Meta `path:"/user/coupon/update/{Id}"method:"post" tags:"用户优惠券"summary:"更新用户优惠券"`
	Id     uint `json:"id" v:"min:1#请选择需要修改的用户优惠券"dc:"用户优惠券Id"`
	UserCouponCommonAddUpdate
}
type UserCouponUpdateRes struct {
	Id uint `json:"id"`
}

// 分页处理
type UserCouponGetListCommonReq struct {
	g.Meta `path:"/user/coupon/list" method:"get" tags:"用户优惠券"summary:"查看用户优惠券"`
	CommonPaginationReq
}
type UserCouponGetListCommonRes struct {
	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}
