package login

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"hellpgf/internal/dao"
	"hellpgf/internal/model"
	"hellpgf/internal/model/entity"
	"hellpgf/internal/service"
	"hellpgf/utility"
)

type sLogin struct{}

func init() {
	service.RegisterLogin(New())
}

func New() *sLogin {
	return &sLogin{}
}

// 执行登录
func (s *sLogin) Login(ctx context.Context, in model.UserLoginInput) error {
	//验证账号密码是否正确
	adminInfo := new(entity.AdminInfo)
	err := dao.AdminInfo.Ctx(ctx).Where("name", in.Name).Scan(&adminInfo)
	if err != nil {
		return err
	}
	if utility.EncryptPassword(in.Password, adminInfo.UserSalt) != adminInfo.Password {
		return gerror.New("账号或者密码不正确")
	}
	if err := service.Session().SetUser(ctx, adminInfo); err != nil {
		return err
	}
	// 自动更新上线
	service.BizCtx().SetUser(ctx, &model.ContextUser{
		Id:      uint(adminInfo.Id),
		Name:    adminInfo.Name,
		IsAdmin: uint8(adminInfo.IsAdmin),
	})
	return nil
}

// 注销
func (s *sLogin) Logout(ctx context.Context) error {
	return service.Session().RemoveUser(ctx)
}
