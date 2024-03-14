// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"hellpgf/internal/dao/internal"
)

// internalCategoryDao is internal type for wrapping internal DAO implements.
type internalCategoryDao = *internal.CategoryDao

// categoryDao is the data access object for table category.
// You can define custom methods on it to extend its functionality as you wish.
type categoryDao struct {
	internalCategoryDao
}

var (
	// Category is globally public accessible object for table category operations.
	Category = categoryDao{
		internal.NewCategoryDao(),
	}
)

// Fill with you ideas below.
