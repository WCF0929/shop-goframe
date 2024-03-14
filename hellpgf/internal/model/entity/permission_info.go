// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PermissionInfo is the golang structure for table permission_info.
type PermissionInfo struct {
	Id        int         `json:"id"        ` //
	Name      string      `json:"name"      ` // 权限名称
	Path      string      `json:"path"      ` // 路径
	CreatedAt *gtime.Time `json:"createdAt" ` //
	UpdatedAt *gtime.Time `json:"updatedAt" ` //
	DeletedAt *gtime.Time `json:"deletedAt" ` //
}
