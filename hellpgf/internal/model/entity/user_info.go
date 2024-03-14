// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UserInfo is the golang structure for table user_info.
type UserInfo struct {
	Id           int         `json:"id"           ` //
	Name         string      `json:"name"         ` // 用户名
	Avatar       string      `json:"avatar"       ` // 头像
	Password     string      `json:"password"     ` //
	UserSalt     string      `json:"userSalt"     ` // 加密盐 生成密码用
	Sex          int         `json:"sex"          ` // 1男 2女
	Status       int         `json:"status"       ` // 1正常 2拉黑冻结
	Sign         string      `json:"sign"         ` // 个性签名
	SecretAnswer string      `json:"secretAnswer" ` // 密保问题的答案
	CreatedAt    *gtime.Time `json:"createdAt"    ` //
	UpdatedAt    *gtime.Time `json:"updatedAt"    ` //
	DeletedAt    *gtime.Time `json:"deletedAt"    ` //
}
