package module

type Module struct {
	Package    string   // 包名
	ImportList []string // 导入名
	Struct     string   // 类名，方法属于这个类
	Method     string   // 方法名
	Note       string   // 方法主食
}
