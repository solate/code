package app

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
	err := New("gen").Start()
	if err != nil {
		t.Error(err.Error())
		return
	}
}
