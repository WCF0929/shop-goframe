// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// EmpDao is the data access object for table emp.
type EmpDao struct {
	table   string     // table is the underlying table name of the DAO.
	group   string     // group is the database configuration group name of current DAO.
	columns EmpColumns // columns contains all the column names of Table for convenient usage.
}

// EmpColumns defines and stores column names for table emp.
type EmpColumns struct {
	Id     string // ID
	DeptId string // 所属部门
	Name   string // 姓名
	Gender string // 性别: 0=男 1=女
	Phone  string // 联系电话
	Email  string // 邮箱
	Avatar string // 照片
}

// empColumns holds the columns for table emp.
var empColumns = EmpColumns{
	Id:     "id",
	DeptId: "dept_id",
	Name:   "name",
	Gender: "gender",
	Phone:  "phone",
	Email:  "email",
	Avatar: "avatar",
}

// NewEmpDao creates and returns a new DAO object for table data access.
func NewEmpDao() *EmpDao {
	return &EmpDao{
		group:   "default",
		table:   "emp",
		columns: empColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *EmpDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *EmpDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *EmpDao) Columns() EmpColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *EmpDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *EmpDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *EmpDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
