package controller

import (
	"github.com/gogf/gf/v2/util/gconv"
	"golang.org/x/net/context"
	"hellpgf/api/frontend"
	"hellpgf/internal/model"
	"hellpgf/internal/service"
)

var Comment = cComment{}

type cComment struct {
}

func (a *cComment) Add(ctx context.Context, req *frontend.AddCommentReq) (res *frontend.AddCommentRes, err error) {
	data := model.AddCommentInput{}
	err = gconv.Struct(req, &data)
	if err != nil {
		return nil, err
	}
	collection, err := service.Comment().AddComment(ctx, data)
	if err != nil {
		return nil, err
	}

	return &frontend.AddCommentRes{Id: collection.Id}, nil
}
func (a *cComment) Delete(ctx context.Context, req *frontend.DeleteCommentReq) (res *frontend.DeleteCommentRes, err error) {
	data := model.DeleteCommentInput{}
	err = gconv.Struct(req, &data)
	if err != nil {
		return nil, err
	}
	collection, err := service.Comment().DeleteComment(ctx, data)
	if err != nil {
		return nil, err
	}
	return &frontend.DeleteCommentRes{Id: collection.Id}, nil
}

// Index article list
func (a *cComment) List(ctx context.Context, req *frontend.ListCommentReq) (res *frontend.ListCommentRes, err error) {

	getListRes, err := service.Comment().GetList(ctx, model.CommentListInput{
		Page: req.Page,
		Size: req.Size,
		Type: req.Type,
	})
	if err != nil {
		return nil, err
	}
	return &frontend.ListCommentRes{
		List:  getListRes.List,
		Page:  getListRes.Page,
		Size:  getListRes.Size,
		Total: getListRes.Total,
	}, nil
}
