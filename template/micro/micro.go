package micro

import "github.com/solate/util/go/gostring"

type Service struct {
	Package    string   // 包名
	ImportList []string // 引用包列表
	StructName string   // 结构体名称
}

// 初始化
func New(p string, i []string, s string) *Service {
	return &Service{
		Package:    p,
		ImportList: i,
		StructName: gostring.Ucfirst(s),
	}
}
