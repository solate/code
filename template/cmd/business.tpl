package app

import (
	{{- range $key, $value := .ImportList}}
    "{{$value}}"
    {{- end}}
)


// 初始化业务组件
func InitBusiness() (err error) {
	// 启用http接口
	go server.New(Config.Port).Start(code.Router)

	return
}
