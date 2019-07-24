package tmpl

import (
	"testing"
)

const TEMP = `



func {{.FucName}}() {
	//生成方法
}


`

type TStruct struct {
	FucName string //名字
}

func TestTemplateParse(t *testing.T) {
	bb := map[string]interface{}{
		"FucName": "code",
	}
	str, err := TemplateParse("test", TEMP, bb)
	if err != nil {
		t.Error(err)
	}

	t.Log(str)

	cc := TStruct{FucName: "msg"}
	str, err = TemplateParse("test", TEMP, cc)
	if err != nil {
		t.Error(err)
	}

	t.Log(str)
}
