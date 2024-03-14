package backend

import (
	"github.com/gogf/gf/v2/frame/g"
	"hellpgf/internal/model/entity"
	"time"
)

/*type LoginIndexReq struct {
	g.Meta `path:"/login" method:"get" summary:"展示登录页面" tags:"登录"`
}
type LoginIndexRes struct {
	g.Meta `mime:"/backend/text/html" type:"string" example:"<html/>"`
}*/

type LoginDoReq struct {
	g.Meta   `path:"/login" method:"post" summary:"执行登录请求" tags:"登录"`
	Name     string `json:"name" v:"required#请输入账号"   dc:"账号"`
	Password string `json:"password" v:"required#请输入密码"   dc:"密码(明文)"`
	//Captcha  string `json:"captcha"  v:"required#请输入验证码" dc:"验证码"`
}

// for jwt
type LoginDoRes struct {
	//User interface{} `json:"user"`
	//todo
	//Referer string `json:"referer" dc:"引导客户端跳转地址"`
	Token  string    `json:"token"`
	Expire time.Time `json:"expire"`
}

// for gtoken
type LoginRes struct {
	Type        string                  `json:"type"`
	Token       string                  `json:"token"`
	ExpireIn    int                     `json:"expireIn"`
	IsAdmin     int                     `json:"isAdmin"`     //是否超管
	RoleIds     string                  `json:"roleIds"`     //角色
	Permissions []entity.PermissionInfo `json:"permissions"` //权限管理
}

type RefreshTokenReq struct {
	g.Meta `path:"/refresh_token" method:"post"`
}

type RefreshTokenRes struct {
	Token  string    `json:"token"`
	Expire time.Time `json:"expire"`
}

type LogoutReq struct {
	g.Meta `path:"/logout" method:"post"`
}

type LogoutRes struct {
}
