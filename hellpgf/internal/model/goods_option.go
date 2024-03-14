package model

import "hellpgf/internal/model/entity"

// GoodsOptionsCreateUpdateBase 创建/修改内容基类
type GoodsOptionsCreateUpdateBase struct {
	GoodsId uint   `json:"goods_id"          v:"required#商品id必填" dc:"商品id"`      //商品id
	PicUrl  string `json:"picUrl"            dc:"图片"        `                    // 图片
	Name    string `json:"name"              dc:"商品操作分类" v:"required#商品操作名称必填" ` // 商品操作名称
	Price   uint   `json:"price"             dc:"价格"    v:"required#价格必填"`       // 价格 单位分
	Stock   uint   `json:"stock"              dc:"库存" `                          // 库存
}

// GoodsOptionsCreateInput  创建内容
type GoodsOptionsCreateInput struct {
	GoodsOptionsCreateUpdateBase
}

// GoodsOptionsCreateOutput 创建内容返回结果
type GoodsOptionsCreateOutput struct {
	Id uint `json:"id"`
}

// GoodsOptionsUpdateInput 修改内容
type GoodsOptionsUpdateInput struct {
	GoodsOptionsCreateUpdateBase
	Id uint
}

// GoodsOptionsGetListInput 获取内容列表
type GoodsOptionsGetListInput struct {
	Page int // 分页号码
	Size int // 分页数量，最大50
}

// GoodsOptionsGetListOutput 查询列表结果
type GoodsOptionsGetListOutput struct {
	List  []GoodsOptionsGetListOutputItem `json:"list" description:"列表"`
	Page  int                             `json:"page" description:"分页码"`
	Size  int                             `json:"size" description:"分页数量"`
	Total int                             `json:"total" description:"数据总数"`
}

// GoodsOptionsSearchInput 搜索列表
type GoodsOptionsSearchInput struct {
	Key            string // 关键字
	Type           string // 内容模型
	GoodsOptionsId uint   // 栏目ID
	Page           int    // 分页号码
	Size           int    // 分页数量，最大50
}

// GoodsOptionsSearchOutput 搜索列表结果
type GoodsOptionsSearchOutput struct {
	List  []GoodsOptionsSearchOutputItem `json:"list"`  // 列表
	Stats map[string]int                 `json:"stats"` // 搜索统计
	Page  uint                           `json:"page"`  // 分页码
	Size  uint                           `json:"size"`  // 分页数量
	Total uint                           `json:"total"` // 数据总数
}

type GoodsOptionsGetListOutputItem struct {
	entity.GoodsOptionsInfo
}

type GoodsOptionsSearchOutputItem struct {
	GoodsOptionsGetListOutputItem
}
