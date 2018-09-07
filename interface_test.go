package generation

import (
	"fmt"
	"strings"
	"testing"
)

func TestGenerateInterfaceCode(t *testing.T) {

	interfaceStr := "Department(*DepartmentReq, *DepartmentReply) error                   //获取科室信息"
	templatePath := "./testdata/virhis.txt"
	writePath := "./export"

	list := strings.Split(interfaceStr, "\n") //按行切分
	for _, v := range list {

		//分割接口，获得接口方法名和注释
		methodName, note := SpliteInterface(v)
		t.Log(methodName, note, "=>=>")

		//放参数
		fileName := methodName
		methodReq := methodName + "Req"
		methodReply := methodName + "Reply"
		methodURL := "URL" + methodName
		reqStr := genReq(methodReq)

		//放参数
		err := GenerateInterfaceCode(templatePath, writePath, fileName,
			note, methodName, methodReq, methodReply, reqStr, methodURL, methodReply)
		if err != nil {
			t.Error(err)
		}

	}

}

func genReq(sub string) string {
	str, err := GenerationReqByFile("./testdata/a.go", sub, "data.Add(\"%s\", %s)")
	if err != nil {
		fmt.Println(err)
	}
	return str
}
