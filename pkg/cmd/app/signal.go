package app

import (
	"github.com/solate/code/pkg/component/logger"
	"github.com/solate/code/pkg/template"
	"path"
)

// 一个服务
type Signal struct {
	pkg string
	*template.Option
}

func NewSignal(packageName string) *Signal {
	o := &template.Option{
		Package: "app", // 固定
		ImportList: []string{
			"github.com/solate/code/pkg/Business/logger",
		},
	}
	return &Signal{packageName, o}
}

func (s *Signal) Start() (err error) {
	logger.Logger.Debug("In Start")
	t := template.New(
		templateSignalPath,
		path.Join(exportDir, s.pkg, "app/signal.go"),
		s.Option,
	)
	return t.Start()
}
