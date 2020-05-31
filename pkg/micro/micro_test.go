package micro

import (
	"github.com/solate/code/pkg/component/logger"
	"testing"
)

func init() {
	if err := logger.Init("debug", ""); err != nil {
		return
	}
}

func TestNewDefault(t *testing.T) {
	err := NewDefault("micro").Start()
	if err != nil {
		t.Error(err.Error())
		return
	}
}

func TestNew(t *testing.T) {
	err := NewDefault("main").Start()
	if err != nil {
		t.Error(err.Error())
		return
	}
}
