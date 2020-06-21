package app

import (
	"github.com/solate/code/pkg/component/logger"
)

// 一个服务
type App struct {
	pkg string
}

func New(pkg string) *App {
	return &App{pkg}
}

func (s *App) Start() (err error) {
	logger.Logger.Debug("In Start")

	// 生成component
	err = NewComponent(s.pkg).Start()
	if err != nil {
		return
	}

	// 生成business
	err = NewBusiness(s.pkg).Start()
	if err != nil {
		return
	}

	// 生成business
	err = NewSignal(s.pkg).Start()
	if err != nil {
		return
	}

	return
}
