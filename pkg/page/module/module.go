package module

import (
	"path"

	"github.com/solate/code/pkg/template"
	"github.com/solate/code/template/module"
)

// 服务参数
type Module struct {
	name string // 模块名称
	*template.Template
}

func New(name string, s *module.Module) *Module {
	t := template.New(
		templateName,
		templateRelativePath,
		path.Join(exportDir, name, name+".go"),
		s,
	)
	return &Module{Template: t}
}

func (s *Module) Start() (err error) {
	return s.Template.Start()
}
