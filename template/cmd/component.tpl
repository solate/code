package app

import (
	"github.com/solate/code/pkg/component/logger"
	"github.com/solate/code/pkg/component/orm"
	"github.com/solate/util/project/configuration"
	"github.com/solate/util/project/errorUtil"

	"github.com/solate/code/cmd/generation/app/config"
)

var (
	Config config.Config
)

func InitComponent() (err error) {
	if err := configuration.LoadConfig(&Config); err != nil {
		return errorUtil.ErrConfig
	}

	if err := logger.Init(Config.Log.Level, Config.Log.Path); err != nil {
		return errorUtil.ErrConfig
	}
	logger.Logger.Debug("init logger end")

	debug := Config.Mode == "debug"
	if err := orm.Init(
		Config.DB.Host,
		Config.DB.Port,
		Config.DB.Name,
		Config.DB.Username,
		Config.DB.Password,
		Config.DB.MaxIdleConns,
		Config.DB.MaxOpenConns,
		Config.DB.ConnMaxLifetime,
		debug,
	); err != nil {
		return errorUtil.ErrConfig
	}
	logger.Logger.Debug("init logger end")

	return
}
