package mainFile

import (
	"github.com/solate/code/pkg/template"
	"github.com/solate/code/template/cmd"
	"path"
)

// 一个服务
type Main struct {
	name               string // 服务名称
	*template.Template        //模板
}

// 初始化
// name: 服务名称
// root: 服务导入根节点
func New(name, root string) *Main {

	s := &cmd.Main{
		ImportList: []string{
			path.Join(root, "/cmd/", name, "/app"),
		},
	}

	t := template.New(
		templateName,
		templateRelativePath,
		path.Join(exportDir, name, "main.go"),
		s,
	)
	return &Main{Template: t}
}

func (s *Main) Start() (err error) {
	return s.Template.Start()
}
