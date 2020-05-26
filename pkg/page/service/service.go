package controller

import (
	"github.com/solate/code/pkg/component/value"
	"github.com/solate/code/pkg/template"
	"github.com/solate/code/template/controller"
	"path"
)

// 服务参数
type Service struct {
	name string // 模块名称
	*template.Template
}

// 初始化
func New(name string, s *controller.Controller) *Service {

	t := &template.Template{
		Name:         templateName,
		TemplatePath: path.Join(value.Root, templateRelativePath),
		ExportPath:   path.Join(value.Export, exportDir, name, name+".go"),
		Structure:    s,
	}
	return &Service{Template: t}
}

// 启动服务
func (s *Service) Start() (err error) {
	if err = s.Template.CheckDir(); err != nil {
		return
	}
	return s.Template.Write()
	return
}
