package models

type Header struct {
	Package string //包名
	Import  string //引入入
}

type Service struct {
}

type Template struct {
	Template string   //模板字符串
	Params   []string //
}
