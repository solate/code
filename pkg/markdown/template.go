package markdown

import (
	"fmt"
	"strings"

	"github.com/solate/code/utils"
	"github.com/solate/util/go/gostring"

	"github.com/solate/code/pkg/markdown/models"
)

const (
	TemplatePackage = "package %s" //包模板
	TemplateImport  = `
import (
	%s
)`
	TemplateHeader = `
package %s

import (
	%s
)

`

	TemplateRouter            = "%s func(*gin.Context) `path:\"%s\"` //%s" //service router
	TemplateMethodDeclaration = "%s: %s,"                                  //service 冒号方法

	TemplateReqBody   = "%s %s `form:\"%s\" json:\"%s\" binding:\"%s\"` //%s" //module req
	TemplateReplyBody = "%s   %s `json:\"%s\"`   //%s"                        //module reply
)

//获取包模板内容
func GetTemplatePackage(pkg string) string {
	return fmt.Sprintf(TemplatePackage, pkg)
}

//获得导入模板内容
func GetTemplateImport(i string) string {
	return fmt.Sprintf(TemplateImport, i)
}

//获得service 路由body
func GetTemplateRouter(list []models.MarkDown) string {

	tmp := make([]string, 0, len(list))
	for _, v := range list {

		tmp = append(tmp, fmt.Sprintf(TemplateRouter,
			GetUriMethodName(v.URL, v.Method), v.URL, v.Note))
	}

	return strings.Join(tmp, "\n")
}

//获得方法列表体字符串
func GetTemplateMethodDeclaration(list []models.MarkDown) string {
	tmp := make([]string, 0, len(list))
	for _, v := range list {
		methodName := utils.GetUriMethodName(v.URL, v.Method)
		tmp = append(tmp, fmt.Sprintf(TemplateMethodDeclaration, methodName, methodName))
	}

	return strings.Join(tmp, "\n")
}

//获得req body内容
func GetTemplateReqBody(list []models.MarkDownReqTable) string {

	tmp := make([]string, 0, len(list))
	for _, v := range list {
		reqForm := gostring.Lcfirst(v.Name)

		isRequired := "omitempty"
		if v.IsRequired == "是" || v.IsRequired == "true" {
			isRequired = "required"
		}

		tmp = append(tmp, fmt.Sprintf(TemplateReqBody,
			v.Name, v.Type, reqForm, reqForm, isRequired, v.Note))
	}

	return strings.Join(tmp, "\n")
}

//获得Response body 内容
func GetTemplateReplyBody(list []models.MarkDownReplyTable) string {

	tmp := make([]string, 0, len(list))
	for _, v := range list {
		reqForm := gostring.Lcfirst(v.Name)

		tmp = append(tmp, fmt.Sprintf(TemplateReplyBody,
			v.Name, v.Type, reqForm, v.Note))
	}

	return strings.Join(tmp, "\n")
}

// url类型, 生成方法名,
// 例如: /register/test => RegisterTest
// path 方法路径
// method 方法类型，get/post
// 方法名 method + path 首字母都大写
func GetUriMethodName(path string, method string) (methodName string) {

	methodName = strings.Title(strings.ToLower(method))
	words := strings.Split(path, "/") //路径以 / 分割
	for _, v := range words {
		methodName += strings.Title(v)
	}
	return
}
