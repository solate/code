package generation

import (
	"io/ioutil"
	"testing"
)

func TestGenerationReq(t *testing.T)  {

	//读取文件
	fileBytes, err := ioutil.ReadFile("testdata/a.go")
	if err != nil {
		t.Error(err)
	}
	//文件字符
	fileStr := string(fileBytes)
	//目标子串
	sub := "DepartmentReq"

	list := GenerationReq(fileStr, sub, "data.Add(\"%s\", %s)")
	t.Log(list)
}

