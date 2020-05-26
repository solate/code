package mainTemplate

import (
	"github.com/solate/code/template/cmd"
	"testing"
)

func TestNew(t *testing.T) {
	s := &cmd.Main{
		ImportList: []string{},
	}

	err := New("gen", s).Start()
	if err != nil {
		t.Error(err.Error())
		return
	}
}
