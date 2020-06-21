package module

import (
	"testing"

	"github.com/solate/code/template/module"
	"github.com/solate/util/go/gostring"
)

func TestNew(t *testing.T) {
	s := &module.Module{
		Package:    "module",
		ImportList: []string{"fmt", "os"},
		Struct:     gostring.Ucfirst("pro"),
		Method:     "List",
		Note:       "获得列表",
	}

	err := New("module", s).Start()
	if err != nil {
		t.Error(err.Error())
		return
	}
}
