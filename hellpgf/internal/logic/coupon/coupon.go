package coupon

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"hellpgf/internal/model/entity"
	"hellpgf/internal/service"

	"hellpgf/internal/dao"
	"hellpgf/internal/model"
)

type sCoupon struct{}

func init() {
	service.RegisterCoupon(New())
}

func New() *sCoupon {
	return &sCoupon{}
}

// Create 创建分类
func (s *sCoupon) Create(ctx context.Context, in model.CouponCreateInput) (out model.CouponCreateOutput, err error) {
	lastInsertID, err := dao.CouponInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return model.CouponCreateOutput{CouponId: uint(int(lastInsertID))}, err
}

// Delete 删除
func (s *sCoupon) Delete(ctx context.Context, id uint) error {
	_, err := dao.CouponInfo.Ctx(ctx).Where(g.Map{
		dao.CouponInfo.Columns().Id: id,
	}).Delete()
	return err
}

// Update 修改
func (s *sCoupon) Update(ctx context.Context, in model.CouponUpdateInput) error {
	_, err := dao.CouponInfo.
		Ctx(ctx).
		Data(in).
		FieldsEx(dao.CouponInfo.Columns().Id).
		Where(dao.CouponInfo.Columns().Id, in.Id).
		Update()
	return err
}

// GetList 查询分类列表
func (s *sCoupon) GetList(ctx context.Context, in model.CouponGetListInput) (out *model.CouponGetListOutput, err error) {
	var ( //m就是model对象
		m = dao.CouponInfo.Ctx(ctx)
	)
	out = &model.CouponGetListOutput{
		Page: in.Page,
		Size: in.Size,
	}
	// 分ye查询
	listModel := m.Page(in.Page, in.Size)
	// 排序方式 按照优惠券价格倒叙排序
	listModel = listModel.OrderDesc(dao.CouponInfo.Columns().Price)

	// 执行查询
	var list []*entity.CouponInfo
	if err := listModel.Scan(&list); err != nil {
		return out, err
	}
	// 没有数据
	if len(list) == 0 {
		return out, nil
	}
	out.Total, err = m.Count()
	if err != nil {
		return out, err
	}
	// Coupon
	if err := listModel.Scan(&out.List); err != nil {
		return out, err
	}
	return
}

// GetAllList 查询全部列表
func (s *sCoupon) GetAllList(ctx context.Context, in model.CouponGetListInput) (out *model.CouponGetListOutput, err error) {
	var ( //m就是model对象
		m = dao.CouponInfo.Ctx(ctx)
	)
	listModel := m
	listModel = m.Page(in.Page, in.Size)
	// 排序方式
	listModel = listModel.OrderDesc(dao.CouponInfo.Columns().Price)
	// 执行查询
	var list []*entity.CouponInfo
	if err := listModel.Scan(&list); err != nil {
		return out, err
	}
	// 没有数据
	if len(list) == 0 {
		return out, nil
	}
	out.Total, err = m.Count()
	if err != nil {
		return out, err
	}
	// Coupon
	if err = listModel.Scan(&out.List); err != nil {
		return out, err
	}
	return
}
