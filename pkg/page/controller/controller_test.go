package controller

import (
	"github.com/solate/code/template/controller"
	"testing"
)

func TestNew(t *testing.T) {
	s := &controller.Controller{
		StructName: "controller",
	}

	err := New("controller", s).Start()
	if err != nil {
		t.Error(err.Error())
		return
	}
}
