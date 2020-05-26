package controller

import (
	"github.com/solate/code/template/controller"
	"testing"
)

func TestNew(t *testing.T) {
	s := &controller.Controller{
		StructName: "service",
	}

	err := New("service", s).Start()
	if err != nil {
		t.Error(err.Error())
		return
	}
}
