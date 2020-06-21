package app

import (
	"github.com/solate/code/pkg/component/logger"
	"github.com/solate/code/pkg/template"
	"path"
)

// 一个服务
type Business struct {
	pkg string
	*template.Option
}

func NewBusiness(packageName string) *Business {
	o := &template.Option{
		Package: "app", // 固定
		ImportList: []string{
			"github.com/solate/code/pkg/Business/logger",
		},
	}
	return &Business{packageName, o}
}

func (s *Business) Start() (err error) {
	logger.Logger.Debug("In Start")
	t := template.New(
		templateBusinessPath,
		path.Join(exportDir, s.pkg, "app/business.go"),
		s.Option,
	)
	return t.Start()
}
