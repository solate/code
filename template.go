package generation

import (
	"fmt"
	"github.com/solate/generation/model"
	"github.com/solate/generation/template"
	"github.com/solate/generation/utils"
	"strings"
)

//获取包模板内容
func GetTemplatePackage(pkg string) string {
	return fmt.Sprintf(template.TemplatePackage, pkg)
}

//获得导入模板内容
func GetTemplateImport(i string) string {
	return fmt.Sprintf(template.TemplateImport, i)
}

//获得service 路由body
func GetTemplateRouter(list []model.MarkDown) string {

	tmp := make([]string, 0, len(list))
	for _, v := range list {

		tmp = append(tmp, fmt.Sprintf(template.TemplateRouter,
			utils.GetUriMethodName(v.URL, v.Method), v.URL, v.Note))
	}

	return strings.Join(tmp, "\n")
}

//获得方法列表体字符串
func GetTemplateMethodDeclaration(list []model.MarkDown) string {
	tmp := make([]string, 0, len(list))
	for _, v := range list {
		methodName := utils.GetUriMethodName(v.URL, v.Method)
		tmp = append(tmp, fmt.Sprintf(template.TemplateMethodDeclaration, methodName, methodName))
	}

	return strings.Join(tmp, "\n")
}

//获得req body内容
func GetTemplateReqBody(list []model.MarkDownReqTable) string {

	tmp := make([]string, 0, len(list))
	for _, v := range list {
		reqForm := utils.Lcfirst(v.Name)

		isRequired := "omitempty"
		if v.IsRequired == "是" || v.IsRequired == "true" {
			isRequired = "required"
		}

		tmp = append(tmp, fmt.Sprintf(template.TemplateReqBody,
			v.Name, v.Type, reqForm, reqForm, isRequired, v.Note))
	}

	return strings.Join(tmp, "\n")
}

//获得Response body 内容
func GetTemplateReplyBody(list []model.MarkDownReplyTable) string {

	tmp := make([]string, 0, len(list))
	for _, v := range list {
		reqForm := utils.Lcfirst(v.Name)

		tmp = append(tmp, fmt.Sprintf(template.TemplateReplyBody,
			v.Name, v.Type, reqForm, v.Note))
	}

	return strings.Join(tmp, "\n")
}
