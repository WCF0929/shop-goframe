// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AdminInfo is the golang structure for table admin_info.
type AdminInfo struct {
	Id        int         `json:"id"        ` //
	Name      string      `json:"name"      ` // 用户名
	Password  string      `json:"password"  ` // 密码
	RoleIds   string      `json:"roleIds"   ` // 角色ids
	CreatedAt *gtime.Time `json:"createdAt" ` //
	UpdatedAt *gtime.Time `json:"updatedAt" ` //
	UserSalt  string      `json:"userSalt"  ` // 加密盐
	IsAdmin   int         `json:"isAdmin"   ` // 是否超级管理员
	DeletedAt *gtime.Time `json:"deletedAt" ` //
}
