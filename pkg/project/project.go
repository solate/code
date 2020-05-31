package project

import (
	"github.com/solate/code/pkg/component/logger"
	"github.com/solate/code/pkg/page/module"
	"github.com/solate/code/pkg/page/service"
)

// 项目配置
type Project struct {
	*service.Service                  // 生成service
	Modules          []*module.Module // 生成module
}

// 初始化
func New() *Project {
	return &Project{}
}

// 启动服务
func (s *Project) Start() (err error) {
	logger.Logger.Debug("In Start")

	// 生成service层
	if err = s.Service.Start(); err != nil {
		return err
	}

	// 多个模块
	for _, v := range s.Modules {
		// 生成module
		if err = v.Start(); err != nil {
			logger.Logger.Error(err.Error())
			continue
		}
	}

	return
}
