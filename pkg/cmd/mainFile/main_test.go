package mainFile

import (
	"github.com/solate/code/pkg/component/value"
	"strings"
	"testing"
)

func TestNew(t *testing.T) {

	list := strings.Split(value.Export, "src/")
	t.Log(list)

	name := "gen"
	root := list[len(list)-1]
	err := New(name, root).Start()
	if err != nil {
		t.Error(err.Error())
		return
	}
}
