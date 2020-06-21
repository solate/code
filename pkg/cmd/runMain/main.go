package runMain

import (
	"github.com/solate/code/pkg/component/logger"
	"github.com/solate/code/pkg/template"
	"path"
)

// 一个服务
type Main struct {
	packageName      string // 包名
	*template.Option        // 模板参数
}

// 初始化
// name: 服务名称
// root: 服务导入根节点
func New(root, name string) *Main {
	o := &template.Option{
		Package: "main",
		ImportList: []string{
			path.Join(root, "/cmd/", name, "/app"),
		},
	}

	return &Main{name, o}
}

func (s *Main) Start() (err error) {
	logger.Logger.Debug("In Start")
	t := template.New(
		templateRelativePath,
		path.Join(exportDir, s.packageName, "main.go"),
		s.Option,
	)
	return t.Start()
}
