package model

import "github.com/solate/util/go/gostring"

type Model struct {
	StructName    string      // 结构体名称
	TableName     string      // 数据库表名，默认为gorm格式
	AttributeList []Attribute // 属性列表
}

type Attribute struct {
	Name        string // 属性名
	Type        string // 属性类型
	Tag         string // tag 样式、
	Description string // 说明
}

// 初始化
func New(s string, list []Attribute) *Model {
	return &Model{
		StructName:    gostring.Ucfirst(s),
		AttributeList: list,
	}
}
