// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"hellpgf/internal/model"
)

type (
	ICategory interface {
		// Create 创建分类
		Create(ctx context.Context, in model.CategoryCreateInput) (out model.CategoryCreateOutput, err error)
		// Delete 删除
		Delete(ctx context.Context, id uint) error
		// Update 修改
		Update(ctx context.Context, in model.CategoryUpdateInput) error
		// GetList 查询分类列表
		GetList(ctx context.Context, in model.CategoryGetListInput) (out *model.CategoryGetListOutput, err error)
		// GetAllList 查询全部列表
		GetAllList(ctx context.Context, in model.CategoryGetListInput) (out *model.CategoryGetListOutput, err error)
	}
)

var (
	localCategory ICategory
)

func Category() ICategory {
	if localCategory == nil {
		panic("implement not found for interface ICategory, forgot register?")
	}
	return localCategory
}

func RegisterCategory(i ICategory) {
	localCategory = i
}
