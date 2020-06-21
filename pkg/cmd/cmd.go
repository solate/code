package cmd

import (
	"github.com/solate/code/pkg/cmd/app"
	"github.com/solate/code/pkg/cmd/config"
	"github.com/solate/code/pkg/cmd/runMain"
)

// 生成一个完整的cmd
type Cmd struct {
	name string // 包名
	root string // main导入根路径
}

// 初始化
func New(root, name string) *Cmd {
	return &Cmd{name, root}
}

// 启动服务
func (s *Cmd) Start() (err error) {

	err = app.New(s.name).Start()
	if err != nil {
		return
	}

	err = runMain.New(s.root, s.name).Start()
	if err != nil {
		return
	}

	err = config.New(s.name).Start()

	return
}
