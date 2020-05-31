package app

import (
	"github.com/solate/code/pkg/template"
	"path"
)

// 一个服务
type Component struct {
	cmdName            string // cmd启动项名称
	*template.Template        //模板
	//*micro.Service // 模板
}

func New(cmdName string, s interface{}) *Component {

	t := template.New(
		templateName,
		templateRelativePath,
		path.Join(exportDir, cmdName, "app/component.go"),
		s,
	)

	return &Component{Template: t}
}

func (s *Component) Start() (err error) {
	return s.Template.Start()
}
