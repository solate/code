package generation

import (
	"fmt"
	"io/ioutil"
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
	//读取文件
	fileBytes, err := ioutil.ReadFile("./testdata/a.go")
	if err != nil {
		fmt.Println("err read file:", err)
	}
	//文件字符
	fileStr := string(fileBytes)

	list := GenerationReq(fileStr, sub, "data.Add(\"%s\", %s)")
	return strings.Join(list, "\n")
}
