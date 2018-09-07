package generation

import (
	"fmt"
	"github.com/solate/generation/utils"
	"strings"
)

func GenerateMarkdownTable(str string, attributeIndex, typeIndex, noteIndex int) (list []string) {
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
