package generation

//
//import (
//	"github.com/solate/generation/config"
//	"github.com/solate/generation/model"
//	"github.com/solate/generation/utils"
//)
//
////生成module层代码
//func GenerateModule(exampleStr string, md model.MarkDown, params ...interface{}) (err error) {
//
//	//获取包名
//	packageName := utils.GetPackageName(exampleStr)
//	//fmt.Println(packageName)
//
//	//获取import body 内容
//	importBody := utils.GetImportBody(exampleStr)
//	//fmt.Println(importBody)
//
//	structName := config.Config.ModuleName
//	url := config.Config.ModuleUri
//	method := md.Method
//	note := config.Config.ModuleNote
//	methodName := utils.GetUriMethodName(url, method)
//
//	templateStr, err := GenerateTemplate(config.Config.ModuleTemplate,
//		packageName,
//		importBody,
//		//Req
//		methodName+"Req",
//		GetTemplateReqBody(md.Request),
//		//reply
//		methodName+"Reply",
//		methodName+"List",
//		//reply list
//		methodName+"List",
//		GetTemplateReplyBody(md.Response),
//		//func
//		note,
//		structName,
//		methodName,
//		methodName+"Req",
//		methodName+"Reply",
//		"//代码逻辑",
//	)
//
//	err = utils.WriteFile(utils.GenerateUnderlineFileName(config.Config.ModuleExport, structName), templateStr)
//
//	return
//}
