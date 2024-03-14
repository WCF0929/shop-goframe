package backend

import "github.com/gogf/gf/v2/frame/g"

type CouponReq struct {
	g.Meta `path:"/coupon/add" method:"post" tags:"优惠券" summary:"添加优惠券"`
	CouponCommonAddUpdate
}

type CouponCommonAddUpdate struct {
	Name       string `json:"name" v:"required#优惠券名称不能为空" dc:"优惠券名称"`
	Price      int    `json:"price" v:"required#价格不能为空"dc:"商品价格"`
	GoodIds    string `json:"good_ids"dc:"可用商品id 多个够好分隔"`
	CategoryId uint   `json:"category_id"dc:"可以的商品分类"`
}

type CouponRes struct {
	CouponId uint `json:"coupon_id"`
}

type CouponDeleteReq struct {
	g.Meta `path:"/coupon/delete"method:"delete"tags:"优惠券"summary:"删除优惠券"`
	Id     uint `v:"min:1#请选择要删除的优惠券"dc:"优惠券id"`
}
type CouponDeleteRes struct {
}
type CouponUpdateReq struct {
	g.Meta `path:"/coupon/update/{Id}"method:"post" tags:"优惠券"summary:"更新优惠券"`
	Id     uint `json:"id" v:"min:1#请选择需要修改的优惠券"dc:"优惠券Id"`
	CouponCommonAddUpdate
}
type CouponUpdateRes struct {
	Id uint `json:"id"`
}

// 分页处理
type CouponGetListCommonReq struct {
	g.Meta `path:"/coupon/list" method:"get" tags:"优惠券"summary:"查看优惠券"`
	Sort   int `json:"sort"in:"query"dc:"排序"`
	CommonPaginationReq
}
type CouponGetListCommonRes struct {
	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}
type CouponGetListAllCommonReq struct {
	g.Meta `path:"/coupon/list/all" method:"get" tags:"优惠券" summary:"优惠券全部分类"`
}
type CouponGetListAllCommonRes struct {
	List  interface{} `json:"list" description:"列表"`
	Total int         `json:"total" description:"数据总数"`
}
