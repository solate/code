package cmd

import (
	"github.com/solate/code/pkg/template"
)

// 一个启动项
type Cmd struct {
	name      string               // 服务名称
	templates []*template.Template //模板
}

func New(name string) *Cmd {
	return &Cmd{name: name}
}

func (s *Cmd) Start() (err error) {

	//main template

	return
}
