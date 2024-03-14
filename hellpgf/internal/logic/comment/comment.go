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

type sComment struct{}

func init() {
	service.RegisterComment(New())
}

func New() *sComment {
	return &sComment{}
}

func (*sComment) AddComment(ctx context.Context, in model.AddCommentInput) (out *model.AddCommentOutput, err error) {
	in.UserId = gconv.Uint(ctx.Value(consts.CtxUserId))
	id, err := dao.CommentInfo.Ctx(ctx).InsertAndGetId(in)
	if err != nil {
		return nil, err
	}

	return &model.AddCommentOutput{Id: uint(id)}, nil
}

// DeleteComment 兼容处理 优先根据评论id删除  评论id为0在根据对象id和type删除
func (*sComment) DeleteComment(ctx context.Context, in model.DeleteCommentInput) (out *model.DeleteCommentOutput, err error) {
	condition := g.Map{
		dao.CommentInfo.Columns().Id:     in.Id,
		dao.CommentInfo.Columns().UserId: ctx.Value(consts.CtxUserId),
	}
	_, err = dao.CommentInfo.Ctx(ctx).Where(condition).Delete()
	if err != nil {
		return nil, err
	}
	return &model.DeleteCommentOutput{Id: gconv.Uint(in.Id)}, nil
}

// GetList 查询内容列表
func (s *sComment) GetList(ctx context.Context, in model.CommentListInput) (out *model.CommentListOutput, err error) {
	var (
		m = dao.CommentInfo.Ctx(ctx)
	)
	out = &model.CommentListOutput{
		Page: in.Page,
		Size: in.Size,
		List: []model.CommentListOutputItem{}, //数据为空时返回数组而不是null
	}
	// 翻页查询
	listModel := m.Page(in.Page, in.Size)

	//条件查询 只查询客户端传入的数据和我们相关的
	if in.Type != 0 {
		listModel = listModel.Where(dao.CommentInfo.Columns().Type, in.Type)
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

// Comment 抽取方法 获得评论数量的方法 for 商品详情&文章详情
func Comment(ctx context.Context, objectId uint, collectionType uint8) (count int, err error) {
	condition := g.Map{
		dao.CommentInfo.Columns().ObjectId: objectId,
		dao.CommentInfo.Columns().Type:     collectionType,
	}
	count, err = dao.CommentInfo.Ctx(ctx).Where(condition).Count()
	if err != nil {
		return 0, err
	}
	return
}

// CheckCollect 抽取方法判断当前用户是否评论
func CheckCollect(ctx context.Context, in model.CheckIsCommentInput) (bool, error) {
	condition := g.Map{
		dao.CommentInfo.Columns().UserId:   ctx.Value(consts.CtxUserId),
		dao.CommentInfo.Columns().ObjectId: in.ObjectId,
		dao.CommentInfo.Columns().Type:     in.Type,
	}
	count, err := dao.CommentInfo.Ctx(ctx).Where(condition).Count()
	if err != nil {
		return false, err
	}
	if count > 0 {
		return true, nil
	} else {
		return false, nil
	}
}
