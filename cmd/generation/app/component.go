package app

import (
	"github.com/solate/code/pkg/component/logger"
	configuration2 "github.com/solate/util/project/configuration"
	"github.com/solate/util/project/errorUtil"

	"github.com/solate/code/cmd/generation/app/config"
)

var (
	Config config.Config
)

func InitComponent() (err error) {
	if err := configuration2.LoadConfig(&Config); err != nil {
		return errorUtil.ErrConfig
	}

	if err := logger.Init(Config.Log.Level, Config.Log.Path); err != nil {
		return errorUtil.ErrConfig
	}
	logger.Logger.Debug("init logger end")

	return
}
