package cmd

import (
	"github.com/solate/code/pkg/component/logger"
	"testing"
)

func init() {
	if err := logger.Init("debug", ""); err != nil {
		return
	}
}

func TestNew(t *testing.T) {
	root := "github.com/solate/code/export/" //生成文件git路径
	err := New(root, "gen").Start()
	if err != nil {
		t.Error(err.Error())
		return
	}
}
