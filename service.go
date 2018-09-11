package generation

import (
	"github.com/solate/generation/config"
	"github.com/solate/generation/model"
	"github.com/solate/generation/utils"
	"strings"
)

//生成service层代码
func GenerateServiceList(exampleStr string, params ...interface{}) (err error) {

	//获取包名
	packageName := utils.GetPackageName(exampleStr)
	//fmt.Println(packageName)

	//获取import body 内容
	importBody := utils.GetImportBody(exampleStr)
	//fmt.Println(importBody)

	structName := config.Config.ServiceName
	url := config.Config.ServiceUri
	method := config.Config.ServiceMethodType
	note := config.Config.ServiceNote
	methodName := utils.GetUriMethodName(url, method)

	// 获取 struct 属性列表
	//file, err := utils.ReadFile("./source/virhis.go")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//str := utils.GetStructBody(file, "ScheduleListReq")
	//fmt.Println(str)

	templateStr, err := GenerateTemplate(config.Config.ServiceTemplate,
		packageName,
		importBody,
		//struct
		structName,
		methodName,
		url,
		note,
		//struct Name()
		structName,
		strings.ToLower(structName),
		//build
		structName,
		structName,
		structName,
		//build return body
		methodName,
		methodName,
		//func
		note,
		methodName,
		methodName+"Reply",
		methodName+"Req",
		methodName,
	)

	err = utils.WriteFile(utils.GenerateUnderlineFileName(config.Config.ServiceExport, structName), templateStr)

	return
}

func GenerateService(exampleStr string, data model.MarkDown) (err error) {

	//获取包名
	packageName := utils.GetPackageName(exampleStr)
	//fmt.Println(packageName)

	//获取import body 内容
	importBody := utils.GetImportBody(exampleStr)
	//fmt.Println(importBody)

	structName := config.Config.ServiceName
	url := config.Config.ServiceUri
	method := config.Config.ServiceMethodType
	note := config.Config.ServiceNote
	methodName := utils.GetUriMethodName(url, method)
	export := config.Config.ServiceExport

	return

}
