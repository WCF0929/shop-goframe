package user_coupon

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"hellpgf/internal/model/entity"
	"hellpgf/internal/service"

	"hellpgf/internal/dao"
	"hellpgf/internal/model"
)

type sUserCoupon struct{}

func init() {
	service.RegisterUserCoupon(New())
}

func New() *sUserCoupon {
	return &sUserCoupon{}
}

// Create 创建分类
func (s *sUserCoupon) Create(ctx context.Context, in model.UserCouponCreateInput) (out model.UserCouponCreateOutput, err error) {
	lastInsertID, err := dao.UserCouponInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return model.UserCouponCreateOutput{Id: uint(lastInsertID)}, err
}

// Delete 删除
func (s *sUserCoupon) Delete(ctx context.Context, id uint) error {
	_, err := dao.UserCouponInfo.Ctx(ctx).Where(g.Map{
		dao.UserCouponInfo.Columns().Id: id,
	}).Delete()
	return err
}

// Update 修改
func (s *sUserCoupon) Update(ctx context.Context, in model.UserCouponUpdateInput) error {
	_, err := dao.UserCouponInfo.
		Ctx(ctx).
		Data(in).
		FieldsEx(dao.UserCouponInfo.Columns().Id).
		Where(dao.UserCouponInfo.Columns().Id, in.Id).
		Update()
	return err
}

// GetList 查询分类列表
func (s *sUserCoupon) GetList(ctx context.Context, in model.UserCouponGetListInput) (out *model.UserCouponGetListOutput, err error) {
	var ( //m就是model对象
		m = dao.UserCouponInfo.Ctx(ctx)
	)
	out = &model.UserCouponGetListOutput{
		Page: in.Page,
		Size: in.Size,
	}
	// 分ye查询
	listModel := m.Page(in.Page, in.Size)
	// 排序方式 按照优惠券价格倒叙排序
	listModel = listModel.OrderDesc(dao.UserCouponInfo.Columns().CouponId)

	// 执行查询
	var list []*entity.UserCouponInfo
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
	// UserCoupon
	if err := listModel.Scan(&out.List); err != nil {
		return out, err
	}
	return
}

// GetAllList 查询全部列表
func (s *sUserCoupon) GetAllList(ctx context.Context, in model.UserCouponGetListInput) (out *model.UserCouponGetListOutput, err error) {
	var ( //m就是model对象
		m = dao.UserCouponInfo.Ctx(ctx)
	)
	listModel := m
	listModel = m.Page(in.Page, in.Size)
	// 排序方式
	listModel = listModel.OrderDesc(dao.UserCouponInfo.Columns().CouponId)
	// 执行查询
	var list []*entity.UserCouponInfo
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
	// UserCoupon
	if err = listModel.Scan(&out.List); err != nil {
		return out, err
	}
	return
}
