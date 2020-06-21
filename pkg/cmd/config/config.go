package config

import (
	"github.com/solate/code/pkg/component/logger"
	"github.com/solate/code/pkg/template"
	"path"
)

// 一个服务
type Config struct {
	pkg string // 包名
	*template.Option
}

func New(pkg string) *Config {

	o := &template.Option{
		Package: "config", // 固定
		ImportList: []string{
			"github.com/solate/code/pkg/Business/logger",
		},
	}
	return &Config{pkg, o}

}

func (s *Config) Start() (err error) {
	logger.Logger.Debug("In Start")
	err = template.New(
		templateConfigPath,
		path.Join(exportDir, s.pkg, "app/config/config.go"),
		s.Option,
	).Start()
	if err != nil {
		return
	}
	err = template.New(
		templateTomlPath,
		path.Join(exportDir, s.pkg, "app/config/config.toml"),
		s.Option,
	).Start()
	if err != nil {
		return
	}

	return
}
