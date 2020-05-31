package micro

import (
	"github.com/solate/code/pkg/component/logger"
	"github.com/solate/code/pkg/template"
	"github.com/solate/util/go/gostring"
	"path"
)

// 一个服务
type Service struct {
	*template.Option
}

// 初始化
func New(o *template.Option) *Service {
	return &Service{o}
}

// 默认参数初始化
func NewDefault(name string) *Service {
	o := &template.Option{
		Package: name,
		ImportList: []string{
			"github.com/solate/code/pkg/component/logger",
			"github.com/solate/code/pkg/template",
			"github.com/solate/util/go/gostring",
		},
		Struct: template.Struct{
			Note:      "微服务",
			Name:      gostring.Ucfirst(name),
			FieldList: nil,
			FuncList:  nil,
		},
	}
	return &Service{o}
}

// 启动服务
func (s *Service) Start() (err error) {
	logger.Logger.Debug("In Start")
	t := template.New(
		templateRelativePath,
		path.Join(exportDir, s.Option.Package, s.Option.Package+".go"),
		s.Option,
	)
	return t.Start()
}
