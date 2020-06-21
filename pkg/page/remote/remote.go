package remote

import (
	"path"

	"github.com/solate/code/pkg/page/remote"
	"github.com/solate/code/pkg/template"
)

// 服务参数
type Remote struct {
	name               string // 服务名称
	*template.Template        //模板
}

// 初始化
func New(name string, s *remote.Remote) *Remote {
	t := template.New(
		templateName,
		templateRelativePath,
		path.Join(exportDir, name, name+".go"),
		s,
	)
	return &Remote{Template: t}
}

// 启动服务
func (s *Remote) Start() (err error) {
	return s.Template.Start()
}
