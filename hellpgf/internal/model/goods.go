package model

import (
	"hellpgf/internal/model/do"
	"hellpgf/internal/model/entity"
)

// GoodsCreateUpdateBase 创建/修改内容基类
type GoodsCreateUpdateBase struct {
	PicUrl           string `json:"picUrl"            dc:"图片"        `                // 图片
	Name             string `json:"name"              dc:"商品分类" v:"required#商品名称必填" ` // 商品名称
	Price            int    `json:"price"             dc:"价格"    v:"required#价格必填"`   // 价格 单位分
	Level1CategoryId int    `json:"level1CategoryId"  dc:"1级分类id"`                    // 1级分类id
	Level2CategoryId int    `json:"level2CategoryId"  dc:"2级分类id"`                    // 2级分类id
	Level3CategoryId int    `json:"level3CategoryId"  dc:"3级分类id"`                    // 3级分类id
	Brand            string `json:"brand"             dc:"品牌"`                        // 品牌
	Stock            int    `json:"stock"             dc:"库存" `                       // 库存
	Sale             int    `json:"sale"              dc:"销量" `                       // 销量
	Tags             string `json:"tags"              dc:"标签" `                       // 标签
	DetailInfo       string `json:"detailInfo"        dc:"商品详情" `
}

// GoodsCreateInput  创建内容
type GoodsCreateInput struct {
	GoodsCreateUpdateBase
}

// GoodsCreateOutput 创建内容返回结果
type GoodsCreateOutput struct {
	Id uint `json:"id"`
}

// GoodsUpdateInput 修改内容
type GoodsUpdateInput struct {
	GoodsCreateUpdateBase
	Id uint
}

// GoodsGetListInput 获取内容列表
type GoodsGetListInput struct {
	Page int // 分页号码
	Size int // 分页数量，最大50
}

// GoodsGetListOutput 查询列表结果
type GoodsGetListOutput struct {
	List  []GoodsGetListOutputItem `json:"list" description:"列表"`
	Page  int                      `json:"page" description:"分页码"`
	Size  int                      `json:"size" description:"分页数量"`
	Total int                      `json:"total" description:"数据总数"`
}

// GoodsSearchInput 搜索列表
type GoodsSearchInput struct {
	Key     string // 关键字
	Type    string // 内容模型
	GoodsId uint   // 栏目ID
	Page    int    // 分页号码
	Size    int    // 分页数量，最大50
}

// GoodsSearchOutput 搜索列表结果
type GoodsSearchOutput struct {
	List  []GoodsSearchOutputItem `json:"list"`  // 列表
	Stats map[string]int          `json:"stats"` // 搜索统计
	Page  uint                    `json:"page"`  // 分页码
	Size  uint                    `json:"size"`  // 分页数量
	Total uint                    `json:"total"` // 数据总数
}

type GoodsGetListOutputItem struct {
	entity.GoodsInfo //todo

}

type GoodsSearchOutputItem struct {
	GoodsGetListOutputItem
}

type GoodsDetailInput struct {
	Id uint
}

type GoodsDetailOutput struct {
	do.GoodsInfo
	Options   []*do.GoodsOptionsInfo `orm:"with:goods_id=id"` //规格sku
	Comments  []*CommentBase         `orm:"with:object_id=id, where:type=1"`
	IsCollect bool                   `json:"is_collect"`
}
