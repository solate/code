package micro

import (
	"github.com/solate/code/pkg/component/value"
	"github.com/solate/code/pkg/template"
	"github.com/solate/code/template/micro"
	"path"
)

// 一个服务
type Service struct {
	name               string // 服务名称
	*template.Template        //模板
	//*micro.Service // 模板
}

func New(name string, s *micro.Service) *Service {
	t := &template.Template{
		Name:         templateName,
		TemplatePath: path.Join(value.Root, templateRelativePath),
		ExportPath:   path.Join(value.Export, exportDir, name, name+".go"),
		Structure:    s,
	}
	return &Service{Template: t}
}

func (s *Service) Start() (err error) {
	if err = s.Template.CheckDir(); err != nil {
		return
	}
	return s.Template.Write()
}
