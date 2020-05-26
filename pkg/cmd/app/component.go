package app

import (
	"github.com/solate/code/pkg/component/value"
	"github.com/solate/code/pkg/template"
	"path"
)

// 一个服务
type Main struct {
	cmdName            string // cmd启动项名称
	*template.Template        //模板
	//*micro.Service // 模板
}

func New(cmdName string, s interface{}) *Main {
	t := &template.Template{
		Name:         templateName,
		TemplatePath: path.Join(value.Root, templateRelativePath),
		ExportPath:   path.Join(value.Export, exportDir, cmdName, "app/component.go"),
		Structure:    s,
	}
	return &Main{Template: t}
}

func (s *Main) Start() (err error) {
	if err = s.Template.CheckDir(); err != nil {
		return
	}
	return s.Template.Write()
}
