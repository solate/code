package micro

import (
	"github.com/solate/code/template/micro"
	"github.com/solate/util/go/gostring"
	"testing"
)

func TestNew(t *testing.T) {
	s := &micro.Service{
		Package:    "service",
		ImportList: []string{"fmt", "os"},
		StructName: gostring.Ucfirst("service"),
	}

	err := New("test", s).Start()
	if err != nil {
		t.Error(err.Error())
		return
	}
}
