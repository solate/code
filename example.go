package generation

//
// 该文件的方法是使用已经写完成的例子来生成代码
import (
	"fmt"
	"github.com/solate/generation/utils"
	"strings"
)

//读取模板文件，生成方法列表字符串
// file 模板路径
func ExampleByFile(file string, params ...interface{}) (str string, err error) {
	//service 模板
	fileStr, err := utils.ReadFile(file)
	if err != nil {
		return
	}
	return Example(fileStr, params...), nil
}

//根据传入模板,参数，生成一个整体的字符串
func Example(example string, params ...interface{}) string {
	return strings.TrimSpace(fmt.Sprintf(example, params...))
}
