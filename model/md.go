package model

type MarkDown struct {
	Note     string          //方法注释
	URL      string          //方法请求路由地址
	Method   string          //方法类型 GET/POST
	Request  []MarkDownTable //请求参数
	Response []MarkDownTable //返回值参数
}

//表格字段
type MarkDownTable struct {
	Name       string //字段名称
	Type       string //字段类型
	Note       string //字段注释
	IsRequired string //是否必填: 是/否, true/false 字符串
}
