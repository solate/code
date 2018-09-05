package utils

import (
	"bufio"
	"io/ioutil"
	"os"
)

//文件是否存在
//fileName 文件名
func IsExistFile(fileName string) (is bool) {
	// 文件不存在则返回error
	_, err := os.Stat(fileName)
	is = os.IsExist(err)
	return
}

//将内容写入文件中，带缓存
//fileName 文件名
//data 写入数据
func WriteFile(fileName, data string) (err error) {
	// 可写方式打开文件, 每次都清空后写入
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		return
	}
	defer file.Close()

	// 为这个文件创建buffered writer
	bufferedWriter := bufio.NewWriter(file)

	// 写字符串到buffer
	_, err = bufferedWriter.WriteString(data)
	if err != nil {
		return
	}

	// 写内存buffer到硬盘
	bufferedWriter.Flush()

	return
}

//读取文件内容
func ReadFile(fileName string) (str string, err error) {

	// 读取文件到byte slice中
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return
	}
	str = string(data)
	return
}

//获取按行读取scanner
func Scanner(fileName string) (scanner *bufio.Scanner, err error) {
	file, err := os.Open(fileName)
	if err != nil {
		return
	}
	scanner = bufio.NewScanner(file)
	return
}

//根据传入的scanner 读取一行
func ReadLine(scanner *bufio.Scanner) (str string, err error) {

	success := scanner.Scan()
	if success == false {
		// 出现错误或者EOF是返回Error
		err = scanner.Err()
		if err != nil {
			return
		}
	}
	return scanner.Text(), nil
}
