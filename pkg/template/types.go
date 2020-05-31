package template

// 构造一个完整的文件，需要的所有的参数, 选择性填写
type Option struct {
	Package    string   // 包名
	File       string   // 文件名， 默认与包名相同
	ImportList []string // 引用包列表
	Struct     Struct   // 文件中结构体，默认一个文件最多就包含一个结构体
	FuncList   []Func   // 文件普通方法
}

// 一个结构体
type Struct struct {
	Note      string  // 注释
	Name      string  // 方法名
	FieldList []Field // 属性列表
	FuncList  []Func  // 结构体方法
}

// 结构体的属性
type Field struct {
	Name string // 属性名
	Type string // 属性类型
	Tag  string // 属性的Tag
	Note string // 属性的注释
}

// 方法
type Func struct {
	Note     string // 注释
	Name     string // 方法名
	Request  string // req 参数
	Response string // 返回值方法
	Content  string // 方法内容
}
