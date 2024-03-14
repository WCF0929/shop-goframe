package collection

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"hellpgf/internal/consts"
	"hellpgf/internal/dao"
	"hellpgf/internal/model"
	"hellpgf/internal/service"
)

type sPraise struct{}

func init() {
	service.RegisterPraise(New())
}

func New() *sPraise {
	return &sPraise{}
}

func (*sPraise) AddPraise(ctx context.Context, in model.AddPraiseInput) (out *model.AddPraiseOutput, err error) {
	in.UserId = gconv.Uint(ctx.Value(consts.CtxUserId))
	id, err := dao.PraiseInfo.Ctx(ctx).InsertAndGetId(in)
	if err != nil {
		return nil, err
	}
	return &model.AddPraiseOutput{Id: uint(id)}, nil
}

// DeletePraise 兼容处理 优先根据点赞id删除  点赞id为0在根据对象id和type删除
func (*sPraise) DeletePraise(ctx context.Context, in model.DeletePraiseInput) (out *model.DeletePraiseOutput, err error) {
	if in.Id != 0 {
		//根据主键删除
		_, err = dao.PraiseInfo.Ctx(ctx).WherePri(in.Id).Delete()
		if err != nil {
			return nil, err
		}
		return &model.DeletePraiseOutput{Id: in.Id}, nil
	} else {
		in.UserId = gconv.Uint(ctx.Value(consts.CtxUserId))
		id, err := dao.PraiseInfo.Ctx(ctx).OmitEmpty().Where(in).Delete() //过滤空值
		if err != nil {
			return &model.DeletePraiseOutput{}, err
		}
		return &model.DeletePraiseOutput{Id: gconv.Uint(id)}, nil
	}
}

// GetList 查询内容列表
func (s *sPraise) GetList(ctx context.Context, in model.PraiseListInput) (out *model.PraiseListOutput, err error) {
	var (
		m = dao.PraiseInfo.Ctx(ctx)
	)
	out = &model.PraiseListOutput{
		Page: in.Page,
		Size: in.Size,
		List: []model.PraiseListOutputItem{}, //数据为空时返回数组而不是null
	}
	// 翻页查询
	listModel := m.Page(in.Page, in.Size)

	//条件查询 只查询客户端传入的数据和我们相关的
	if in.Type != 0 {
		listModel = listModel.Where(dao.PraiseInfo.Columns().Type, in.Type)
	}
	//优化：优先查询count而不是直接查sql结果赋值到结构体中
	out.Total, err = listModel.Count()
	if err != nil {
		return out, err
	}
	if out.Total == 0 {
		return out, err
	}
	if in.Type == consts.CollectionTypeGoods {
		err = listModel.With(model.GoodsItem{}).Scan(&out.List)
		if err != nil {
			return out, err
		} else if in.Type == consts.CollectionTypeArticle { //如果只查询商品会减少一次sql查询article
			err = listModel.With(model.ArticleItem{}).Scan(&out.List)
		}
	} else {
		if err := listModel.WithAll().Scan(&out.List); err != nil {
			return out, err
		}
	}
	return
}

// Praise 抽取方法 获得点赞数量的方法 for 商品详情&文章详情
func Praise(ctx context.Context, objectId uint, collectionType uint8) (count int, err error) {
	condition := g.Map{
		dao.PraiseInfo.Columns().ObjectId: objectId,
		dao.PraiseInfo.Columns().Type:     collectionType,
	}
	count, err = dao.PraiseInfo.Ctx(ctx).Where(condition).Count()
	if err != nil {
		return 0, err
	}
	return
}

// CheckCollect 抽取方法判断当前用户是否点赞
func CheckCollect(ctx context.Context, in model.CheckIsPraiseInput) (bool, error) {
	condition := g.Map{
		dao.PraiseInfo.Columns().UserId:   ctx.Value(consts.CtxUserId),
		dao.PraiseInfo.Columns().ObjectId: in.ObjectId,
		dao.PraiseInfo.Columns().Type:     in.Type,
	}
	count, err := dao.PraiseInfo.Ctx(ctx).Where(condition).Count()
	if err != nil {
		return false, err
	}
	if count > 0 {
		return true, nil
	} else {
		return false, nil
	}
}
