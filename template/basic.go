package template

const (
	TemplatePackage = "package %s" //包模板
	TemplateImport  = `
import (
	%s
)`

	TemplateRouter            = "%s func(*gin.Context) `path:\"%s\"` //%s" //service router
	TemplateMethodDeclaration = "%s: %s,"                                  //service 冒号方法

	TemplateReqBody   = "%s %s `form:\"%s\" json:\"%s\" binding:\"%s\"` //%s" //module req
	TemplateReplyBody = "%s   %s `json:\"%s\"`   //%s"                        //module reply
)
