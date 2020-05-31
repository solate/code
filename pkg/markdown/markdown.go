package markdown

import (
	"fmt"
	"strings"

	"github.com/solate/code/pkg/markdown/models"
	"github.com/solate/util/go/gofile"
	"github.com/solate/util/go/gostring"
)

type Markdown struct {
	path string //文件路径
}

func New(path string) *Markdown {
	return &Markdown{path: path}
}

// 启动markdown
func (s *Markdown) Start() (err error) {

	// TODO 解析完整的markdown
	return

}

//传入文件，解析markdown
func (s *Markdown) ParseMarkdownByFile() (list []models.MarkDown, err error) {
	markdown, err := gofile.ReadFile(s.path)
	if err != nil {
		return
	}

	list = s.ParseMarkdown(markdown)
	return
}

//解析书写的需求文档, 获得需要的参数
//默认文档完全符合规范
func (s *Markdown) ParseMarkdown(str string) (list []models.MarkDown) {
	//将整个付给子串变量, 后面部分, 跟着使用子串进行切割
	blocks := SplitBlock(str) //默认使用### 级别是每个文档接口
	list = make([]models.MarkDown, 0, len(blocks))
	for _, block := range blocks {
		attrs := SplitBlockAttr(block)
		var md models.MarkDown
		for _, attr := range attrs {

			if strings.Contains(attr, "Note") {
				md.Note = GetAttrBody(attr, "Note") //获得注释
			}
			if strings.Contains(attr, "URL") {
				md.URL = GetAttrBody(attr, "URL") //获得路由
			}
			if strings.Contains(attr, "Method") {
				md.Method = GetAttrBody(attr, "Method")
			}
			if strings.Contains(attr, "Request") {
				md.Request = GetAttrReqTable(attr) //获得Req
			}
			if strings.Contains(attr, "Response") {
				md.Response = GetAttrReplyTable(attr) //获得Reply
			}
			if strings.Contains(attr, "__list__") {
				md.ResponseType = true
			}
		}

		list = append(list, md)
	}

	return
}

//分割 markdown写的文档块, 获得每个接口整体字符串列表
// source 来源字符串
// tag 分割标记
func SplitBlock(source string) (list []string) {
	return gostring.SplitByTag(source, "###")
}

//分割接口块, 获得每个接口的参数
// note: 注释， URL: 接口路由, Request: 请求参数, Response: 返回参数
func SplitBlockAttr(source string) (list []string) {
	return gostring.SplitByTag(source, "*")
}

//获取tag后指定部分
func GetAttrBody(source, tag string) (str string) {
	index := strings.Index(source, tag)
	str = strings.TrimSpace(source[index+len(tag):]) //+4是为了不包含tag
	return
}

func GetAttrReqTable(source string) (list []models.MarkDownReqTable) {
	begin := strings.Index(source, "|")
	end := strings.LastIndex(source, "|")

	//获取表格
	source = source[begin : end+1] //包含最后一个 |

	strSlice := strings.Split(source, "\n") //按行读取
	strSlice = strSlice[2:]                 //跳过表格表头, 也就是第一行第二行
	list = make([]models.MarkDownReqTable, 0, len(strSlice))
	for _, v := range strSlice {
		//if k == 0 || k == 1 { //跳过表格表头, 也就是第一行第二行
		//	continue
		//}
		value := strings.Split(v, "|")
		if len(value) != 7 { //如果表格不对，直接return
			return
		}

		list = append(list, models.MarkDownReqTable{
			Name:       strings.Title(gostring.UnderlineToCamel(strings.TrimSpace(value[1]))),
			Type:       strings.TrimSpace(value[2]),
			Note:       strings.TrimSpace(value[3]),
			IsRequired: strings.TrimSpace(value[4]),
		})

	}

	return
}

func GetAttrReplyTable(source string) (list []models.MarkDownReplyTable) {

	begin := strings.Index(source, "|")
	end := strings.LastIndex(source, "|")

	//获取表格
	source = source[begin : end+1] //包含最后一个 |

	strSlice := strings.Split(source, "\n") //按行读取
	strSlice = strSlice[2:]                 //跳过表格表头, 也就是第一行第二行
	list = make([]models.MarkDownReplyTable, 0, len(strSlice))
	for _, v := range strSlice {
		//if k == 0 || k == 1 { //跳过表格表头, 也就是第一行第二行
		//	continue
		//}
		value := strings.Split(v, "|")
		if len(value) != 7 { //如果表格不对，直接return
			return
		}

		list = append(list, models.MarkDownReplyTable{
			Name: strings.Title(gostring.UnderlineToCamel(strings.TrimSpace(value[1]))),
			Type: strings.TrimSpace(value[2]),
			Note: strings.TrimSpace(value[3]),
		})

	}

	return
}

func MarkdownTable(str string, attributeIndex, typeIndex, noteIndex int) (list []string) {
	strSlice := strings.Split(str, "\n") //按行读取

	list = make([]string, 0, len(strSlice))
	for _, v := range strSlice {
		value := strings.Split(v, "|")
		attribute := fmt.Sprintf("  %s  %s //%s",
			strings.Title(gostring.UnderlineToCamel(value[attributeIndex])),
			value[typeIndex],
			strings.TrimSpace(value[noteIndex]),
		)
		list = append(list, attribute)
	}
	return
}
