package user_coupon

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"hellpgf/internal/model/entity"
	"hellpgf/internal/service"

	"hellpgf/internal/dao"
	"hellpgf/internal/model"
)

type sGoodsOptions struct{}

func init() {
	service.RegisterGoodsOptions(New())
}

func New() *sGoodsOptions {
	return &sGoodsOptions{}
}

// Create 创建分类
func (s *sGoodsOptions) Create(ctx context.Context, in model.GoodsOptionsCreateInput) (out model.GoodsOptionsCreateOutput, err error) {
	lastInsertID, err := dao.GoodsOptionsInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return model.GoodsOptionsCreateOutput{Id: uint(lastInsertID)}, err
}

// Delete 删除
func (s *sGoodsOptions) Delete(ctx context.Context, id uint) error {
	_, err := dao.GoodsOptionsInfo.Ctx(ctx).Where(g.Map{
		dao.GoodsOptionsInfo.Columns().Id: id,
	}).Delete()
	return err
}

// Update 修改
func (s *sGoodsOptions) Update(ctx context.Context, in model.GoodsOptionsUpdateInput) error {
	_, err := dao.GoodsOptionsInfo.
		Ctx(ctx).
		Data(in).
		FieldsEx(dao.GoodsOptionsInfo.Columns().Id).
		Where(dao.GoodsOptionsInfo.Columns().Id, in.Id).
		Update()
	return err
}

// GetList 查询分类列表
func (s *sGoodsOptions) GetList(ctx context.Context, in model.GoodsOptionsGetListInput) (out *model.GoodsOptionsGetListOutput, err error) {
	var ( //m就是model对象
		m = dao.GoodsOptionsInfo.Ctx(ctx)
	)
	out = &model.GoodsOptionsGetListOutput{
		Page: in.Page,
		Size: in.Size,
	}
	// 分ye查询
	listModel := m.Page(in.Page, in.Size)
	// 执行查询
	var list []*entity.GoodsOptionsInfo
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
	// GoodsOptions
	if err := listModel.Scan(&out.List); err != nil {
		return out, err
	}
	return
}

// GetAllList 查询全部列表
func (s *sGoodsOptions) GetAllList(ctx context.Context, in model.GoodsOptionsGetListInput) (out *model.GoodsOptionsGetListOutput, err error) {
	var ( //m就是model对象
		m = dao.GoodsOptionsInfo.Ctx(ctx)
	)
	listModel := m
	listModel = m.Page(in.Page, in.Size)

	// 执行查询
	var list []*entity.GoodsOptionsInfo
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
	// GoodsOptions
	if err = listModel.Scan(&out.List); err != nil {
		return out, err
	}
	return
}
