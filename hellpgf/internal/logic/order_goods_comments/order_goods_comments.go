package order_goods_comments

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"hellpgf/internal/dao"
	"hellpgf/internal/model"
	"hellpgf/internal/service"
)

type sOrderGoodsComments struct {
}

func init() {
	service.RegisterOrderGoodsComments(New())
}
func New() *sOrderGoodsComments {
	return &sOrderGoodsComments{}
}

func (s *sOrderGoodsComments) Add(ctx context.Context, in model.AddOrderGoodsCommentsInput) (out model.AddOrderGoodsCommentsOutput, err error) {
	id, err := dao.OrderGoodsInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return model.AddOrderGoodsCommentsOutput{}, err
	}
	return model.AddOrderGoodsCommentsOutput{
		Id: gconv.Uint(id),
	}, nil

}

func (s *sOrderGoodsComments) Delete(ctx context.Context, in model.DelOrderGoodsCommentsInput) (out model.DelOrderGoodsCommentsOutput, err error) {
	_, err = dao.OrderGoodsInfo.Ctx(ctx).WherePri(in.Id).Delete()
	if err != nil {
		return model.DelOrderGoodsCommentsOutput{}, err
	}
	return model.DelOrderGoodsCommentsOutput{Id: in.Id}, nil
}
