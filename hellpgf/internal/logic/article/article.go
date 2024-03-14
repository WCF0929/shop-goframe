package user_coupon

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"hellpgf/internal/model/entity"
	"hellpgf/internal/service"

	"hellpgf/internal/dao"
	"hellpgf/internal/model"
)

type sArticle struct{}

func init() {
	service.RegisterArticle(New())
}

func New() *sArticle {
	return &sArticle{}
}

// Create 创建分类
func (s *sArticle) Create(ctx context.Context, in model.ArticleCreateInput) (out model.ArticleCreateOutput, err error) {
	lastInsertID, err := dao.ArticleInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return model.ArticleCreateOutput{Id: uint(lastInsertID)}, err
}

// Delete 删除
func (s *sArticle) Delete(ctx context.Context, id uint) error {
	_, err := dao.ArticleInfo.Ctx(ctx).Where(g.Map{
		dao.ArticleInfo.Columns().Id: id,
	}).Delete()
	return err
}

// Update 修改
func (s *sArticle) Update(ctx context.Context, in model.ArticleUpdateInput) error {
	_, err := dao.ArticleInfo.
		Ctx(ctx).
		Data(in).
		FieldsEx(dao.ArticleInfo.Columns().Id).
		Where(dao.ArticleInfo.Columns().Id, in.Id).
		Update()
	return err
}

// GetList 查询分类列表
func (s *sArticle) GetList(ctx context.Context, in model.ArticleGetListInput) (out *model.ArticleGetListOutput, err error) {
	var ( //m就是model对象
		m = dao.ArticleInfo.Ctx(ctx)
	)
	out = &model.ArticleGetListOutput{
		Page: in.Page,
		Size: in.Size,
	}
	// 分ye查询
	listModel := m.Page(in.Page, in.Size)
	// 执行查询
	var list []*entity.ArticleInfo
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
	// Article
	if err := listModel.Scan(&out.List); err != nil {
		return out, err
	}
	return
}

// GetAllList 查询全部列表
func (s *sArticle) GetAllList(ctx context.Context, in model.ArticleGetListInput) (out *model.ArticleGetListOutput, err error) {
	var ( //m就是model对象
		m = dao.ArticleInfo.Ctx(ctx)
	)
	listModel := m
	listModel = m.Page(in.Page, in.Size)

	// 执行查询
	var list []*entity.ArticleInfo
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
	// Article
	if err = listModel.Scan(&out.List); err != nil {
		return out, err
	}
	return
}
