package app

import (
	"github.com/solate/code/pkg/code"
	"github.com/solate/util/gin/server"
)

// 初始化业务组件
func InitBusiness() (err error) {
	// 启用http接口
	go server.New(Config.Port).Start(code.Router)

	return
}
