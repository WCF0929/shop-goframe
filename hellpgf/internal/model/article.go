package model

import "hellpgf/internal/model/entity"

// ArticleCreateUpdateBase 创建/修改内容基类
type ArticleCreateUpdateBase struct {
	UserId  int
	Title   string
	Desc    string
	PicUrl  string
	IsAdmin uint
	Praise  int
	Detail  string
}

// ArticleCreateInput  创建内容
type ArticleCreateInput struct {
	ArticleCreateUpdateBase
}

// ArticleCreateOutput 创建内容返回结果
type ArticleCreateOutput struct {
	Id uint `json:"id"`
}

// ArticleUpdateInput 修改内容
type ArticleUpdateInput struct {
	ArticleCreateUpdateBase
	Id uint
}

// ArticleGetListInput 获取内容列表
type ArticleGetListInput struct {
	Page int // 分页号码
	Size int // 分页数量，最大50
}

// ArticleGetListOutput 查询列表结果
type ArticleGetListOutput struct {
	List  []ArticleGetListOutputItem `json:"list" description:"列表"`
	Page  int                        `json:"page" description:"分页码"`
	Size  int                        `json:"size" description:"分页数量"`
	Total int                        `json:"total" description:"数据总数"`
}

// ArticleSearchInput 搜索列表
type ArticleSearchInput struct {
	Key       string // 关键字
	Type      string // 内容模型
	ArticleId uint   // 栏目ID
	Page      int    // 分页号码
	Size      int    // 分页数量，最大50
}

// ArticleSearchOutput 搜索列表结果
type ArticleSearchOutput struct {
	List  []ArticleSearchOutputItem `json:"list"`  // 列表
	Stats map[string]int            `json:"stats"` // 搜索统计
	Page  uint                      `json:"page"`  // 分页码
	Size  uint                      `json:"size"`  // 分页数量
	Total uint                      `json:"total"` // 数据总数
}

type ArticleGetListOutputItem struct {
	entity.ArticleInfo
}

type ArticleSearchOutputItem struct {
	ArticleGetListOutputItem
}
