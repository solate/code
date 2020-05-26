package model

import (
	"github.com/solate/code/pkg/component/value"
	"github.com/solate/code/pkg/template"
	"github.com/solate/code/template/model"
	"github.com/solate/util/go/gostring"
	"path"
)

// 服务参数
type Model struct {
	name      string // 模块名称
	tableName string // 数据库名称，默认为gorm定义方式，下划线连接后面加s
	*template.Template
}

// 初始化
func New(name string, s *model.Model) *Model {

	// 默认表名 TODO 这里到时候构造一个方法专门处理这个结构体
	if s.TableName == "" {
		s.TableName = gostring.CamelToUnderline(name) + "s"
	}

	// 默认tag

	t := &template.Template{
		Name:         templateName,
		TemplatePath: path.Join(value.Root, templateRelativePath),
		ExportPath:   path.Join(value.Export, exportDir, name, name+".go"),
		Structure:    s,
	}
	return &Model{Template: t}
}

// 启动服务
func (s *Model) Start() (err error) {
	if err = s.Template.CheckDir(); err != nil {
		return
	}
	return s.Template.Write()
	return
}
