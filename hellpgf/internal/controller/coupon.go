package controller

import (
	"golang.org/x/net/context"
	"hellpgf/api/backend"
	"hellpgf/internal/model"
	"hellpgf/internal/service"
)

// Coupon 内容管理
var Coupon = cCoupon{}

type cCoupon struct{}

func (a *cCoupon) Create(ctx context.Context, req *backend.CouponReq) (res *backend.CouponRes, err error) {
	out, err := service.Coupon().Create(ctx, model.CouponCreateInput{
		CouponCreateUpdateBase: model.CouponCreateUpdateBase{
			Name:       req.Name,
			Price:      req.Price,
			GoodsIds:   req.GoodIds,
			CategoryId: req.CategoryId,
		},
	})
	if err != nil {
		return nil, err
	}
	return &backend.CouponRes{CouponId: out.CouponId}, nil
}
func (a *cCoupon) Delete(ctx context.Context, req *backend.CouponDeleteReq) (res *backend.CouponDeleteRes, err error) {
	err = service.Coupon().Delete(ctx, req.Id)
	return
}
func (a *cCoupon) Update(ctx context.Context, req *backend.CouponUpdateReq) (res *backend.CouponUpdateRes, err error) {
	err = service.Coupon().Update(ctx, model.CouponUpdateInput{
		Id: req.Id,
		CouponCreateUpdateBase: model.CouponCreateUpdateBase{
			Name:       req.Name,
			Price:      req.Price,
			GoodsIds:   req.GoodIds,
			CategoryId: req.CategoryId,
		},
	})
	return &backend.CouponUpdateRes{Id: req.Id}, nil
}

// Index article list
func (a *cCoupon) List(ctx context.Context, req *backend.CouponGetListCommonReq) (res *backend.CouponGetListCommonRes, err error) {

	getListRes, err := service.Coupon().GetList(ctx, model.CouponGetListInput{
		Page: req.Page,
		Size: req.Size,
		Sort: req.Sort,
	})
	if err != nil {
		return nil, err
	}
	return &backend.CouponGetListCommonRes{
		List:  getListRes.List,
		Page:  getListRes.Page,
		Size:  getListRes.Size,
		Total: getListRes.Total,
	}, nil
} // Index article list
func (a *cCoupon) ListAll(ctx context.Context, req *backend.CouponGetListAllCommonReq) (res *backend.CouponGetListAllCommonRes, err error) {

	getListRes, err := service.Coupon().GetList(ctx, model.CouponGetListInput{})
	if err != nil {
		return nil, err
	}
	return &backend.CouponGetListAllCommonRes{
		List:  getListRes.List,
		Total: getListRes.Total,
	}, nil
}
