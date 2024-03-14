package user_coupon

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"hellpgf/internal/consts"
	"hellpgf/internal/logic/comment"
	"hellpgf/internal/model/entity"
	"hellpgf/internal/service"

	"hellpgf/internal/dao"
	"hellpgf/internal/model"
)

type sGoods struct{}

func init() {
	service.RegisterGoods(New())
}

func New() *sGoods {
	return &sGoods{}
}

// Create 创建分类
func (s *sGoods) Create(ctx context.Context, in model.GoodsCreateInput) (out model.GoodsCreateOutput, err error) {
	lastInsertID, err := dao.GoodsInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return model.GoodsCreateOutput{Id: uint(lastInsertID)}, err
}

// Delete 删除
func (s *sGoods) Delete(ctx context.Context, id uint) error {
	_, err := dao.GoodsInfo.Ctx(ctx).Where(g.Map{
		dao.GoodsInfo.Columns().Id: id,
	}).Delete()
	return err
}

// Update 修改
func (s *sGoods) Update(ctx context.Context, in model.GoodsUpdateInput) error {
	_, err := dao.GoodsInfo.
		Ctx(ctx).
		Data(in).
		FieldsEx(dao.GoodsInfo.Columns().Id).
		Where(dao.GoodsInfo.Columns().Id, in.Id).
		Update()
	return err
}

// GetList 查询分类列表
func (s *sGoods) GetList(ctx context.Context, in model.GoodsGetListInput) (out *model.GoodsGetListOutput, err error) {
	var ( //m就是model对象
		m = dao.GoodsInfo.Ctx(ctx)
	)
	out = &model.GoodsGetListOutput{
		Page: in.Page,
		Size: in.Size,
	}
	// 分ye查询
	listModel := m.Page(in.Page, in.Size)
	// 执行查询
	var list []*entity.GoodsInfo
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
	// Goods
	if err := listModel.Scan(&out.List); err != nil {
		return out, err
	}
	return
}

// GetAllList 查询全部列表
func (s *sGoods) GetAllList(ctx context.Context, in model.GoodsGetListInput) (out *model.GoodsGetListOutput, err error) {
	var ( //m就是model对象
		m = dao.GoodsInfo.Ctx(ctx)
	)
	listModel := m
	listModel = m.Page(in.Page, in.Size)

	// 执行查询
	var list []*entity.GoodsInfo
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
	// Goods
	if err = listModel.Scan(&out.List); err != nil {
		return out, err
	}
	return
}

// 商品详情
func (*sGoods) Detail(ctx context.Context, in model.GoodsDetailInput) (out model.GoodsDetailOutput, err error) {
	err = dao.GoodsInfo.Ctx(ctx).WithAll().WherePri(in.Id).Scan(&out)
	if err != nil {
		return model.GoodsDetailOutput{}, err
	}
	out.IsCollect, err = collection.CheckCollect(ctx, model.CheckIsCommentInput{
		UserId:   gconv.Uint(ctx.Value(consts.CtxUserId)),
		ObjectId: in.Id,
		Type:     consts.CollectionTypeGoods,
	})
	return
}
