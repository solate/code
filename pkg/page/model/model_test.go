package model

import (
	"github.com/solate/code/template/model"
	"github.com/solate/util/go/gostring"
	"testing"
)

func TestNew(t *testing.T) {
	s := &model.Model{
		StructName: gostring.Ucfirst("model"),
		AttributeList: []model.Attribute{
			model.Attribute{Name: "Name", Type: "string", Tag: `json:"name"`, Description: "名称"},
			model.Attribute{Name: "Age", Type: "int", Tag: `json:"name"`, Description: "名称"},
		},
	}

	err := New("model", s).Start()
	if err != nil {
		t.Error(err.Error())
		return
	}
}
