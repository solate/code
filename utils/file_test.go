package utils

import "testing"

//写内容测试
func TestWriteFile(t *testing.T) {

	err := WriteFile("../testdata/test.txt", "hello world!")
	if err != nil {
		t.Error(err)
	}
	t.Log("success")
}
