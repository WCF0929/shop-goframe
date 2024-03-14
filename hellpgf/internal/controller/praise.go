package controller

import (
	"github.com/gogf/gf/v2/util/gconv"
	"golang.org/x/net/context"
	"hellpgf/api/frontend"
	"hellpgf/internal/model"
	"hellpgf/internal/service"
)

var Praise = cPraise{}

type cPraise struct {
}

func (a *cPraise) Add(ctx context.Context, req *frontend.AddPraiseReq) (res *frontend.AddPraiseRes, err error) {
	data := model.AddPraiseInput{}
	err = gconv.Struct(req, &data)
	if err != nil {
		return nil, err
	}
	collection, err := service.Praise().AddPraise(ctx, data)
	if err != nil {
		return nil, err
	}

	return &frontend.AddPraiseRes{Id: collection.Id}, nil
}
func (a *cPraise) Delete(ctx context.Context, req *frontend.DeletePraiseReq) (res *frontend.DeletePraiseRes, err error) {
	data := model.DeletePraiseInput{}
	err = gconv.Struct(req, &data)
	if err != nil {
		return nil, err
	}
	collection, err := service.Praise().DeletePraise(ctx, data)
	if err != nil {
		return nil, err
	}
	return &frontend.DeletePraiseRes{Id: collection.Id}, nil
}

// Index article list
func (a *cPraise) List(ctx context.Context, req *frontend.ListPraiseReq) (res *frontend.ListPraiseRes, err error) {

	getListRes, err := service.Praise().GetList(ctx, model.PraiseListInput{
		Page: req.Page,
		Size: req.Size,
		Type: req.Type,
	})
	if err != nil {
		return nil, err
	}
	return &frontend.ListPraiseRes{
		List:  getListRes.List,
		Page:  getListRes.Page,
		Size:  getListRes.Size,
		Total: getListRes.Total,
	}, nil
}
