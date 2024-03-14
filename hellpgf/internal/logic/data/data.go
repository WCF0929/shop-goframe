package data

import (
	"context"
	"github.com/gogf/gf/v2/os/gtime"
	"hellpgf/internal/dao"
	"hellpgf/internal/model"
	"hellpgf/internal/service"
	"hellpgf/utility"
)

type sData struct {
}

// 注册服务
func init() {
	service.RegisterData(New())
}

func New() *sData {
	return &sData{}
}

func (s *sData) DataHead(ctx context.Context) (out *model.DataHeadOutout, err error) {
	return &model.DataHeadOutout{
		TodayOrderCount: todayOrderCount(ctx),
		DAU:             utility.RandInt(1000),
		ConversionRate:  utility.RandInt(50),
	}, err
}

// 查询今天的订单总数
func todayOrderCount(ctx context.Context) (count int) {
	count, err := dao.OrderInfo.Ctx(ctx).WhereBetween(dao.OrderInfo.Columns().CreatedAt, gtime.Now().StartOfDay(), gtime.Now().EndOfDay()).Count(dao.OrderInfo.Columns().Id)
	if err != nil {
		return -1
	}
	return count
}
