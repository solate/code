package app

import (
	"github.com/solate/code/pkg/component/logger"
	"github.com/solate/code/pkg/template"
	"path"
)

// 一个服务
type Component struct {
	pkg string //包名
	*template.Option
}

func NewComponent(packageName string) *Component {

	o := &template.Option{
		Package: "app", // 固定
		ImportList: []string{
			"github.com/solate/code/pkg/component/logger",
			"github.com/solate/code/pkg/template",
			"github.com/solate/util/go/gostring",
		},
	}

	return &Component{packageName, o}
}

func (s *Component) Start() (err error) {
	logger.Logger.Debug("In Start")
	t := template.New(
		templateComponentPath,
		path.Join(exportDir, s.pkg, "app/component.go"),
		s.Option,
	)
	return t.Start()
}
