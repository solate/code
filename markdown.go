package generation

import (
	"fmt"
	"github.com/solate/generation/utils"
	"regexp"
	"strings"
)

//解析书写的需求文档, 获得需要的参数
//默认文档完全符合规范
func ParseMarkdown(md string) (err error) {

	//将整个付给子串变量, 后面部分, 跟着使用子串进行切割
	SplitMarkdown(md)

	return
}

//切割
func SplitMarkdown(source string) (err error) {

	r, err := regexp.Compile(`###[/s/S]*Note(\w*)`)
	if err != nil {
		return
	}

	list := r.FindAllStringSubmatch(source, -1)

	fmt.Println(list)

	return
}

func MarkdownTable(str string, attributeIndex, typeIndex, noteIndex int) (list []string) {
	strSlice := strings.Split(str, "\n") //按行读取

	list = make([]string, 0, len(strSlice))
	for _, v := range strSlice {
		value := strings.Split(v, "|")
		attribute := fmt.Sprintf("  %s  %s //%s",
			strings.Title(utils.UnderlineToCamel(value[attributeIndex])),
			value[typeIndex],
			strings.TrimSpace(value[noteIndex]),
		)
		list = append(list, attribute)
	}
	return
}
