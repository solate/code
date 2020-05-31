package models

type MarkDown struct {
	Note         string               //方法注释
	URL          string               //方法请求路由地址
	Method       string               //方法类型 GET/POST
	Request      []MarkDownReqTable   //请求参数
	Response     []MarkDownReplyTable //返回值参数
	ResponseType bool                 //返回值Response data 是 list/single, 单层还是多层结构
}

//表格字段
type MarkDownReqTable struct {
	Name       string //字段名称
	Type       string //字段类型
	Note       string //字段注释
	IsRequired string //是否必填: 是/否, true/false 字符串
}

type MarkDownReplyTable struct {
	Name string //字段名称
	Type string //字段类型
	Note string //字段注释

}
