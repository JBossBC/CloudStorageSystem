package types

import (
	"testing"
)

func TestFileFactory_CreateFile(t *testing.T) {
	file, err := GetFileFactory().CreateFile("xxx:///D:/CloudStorage/ProjectDesign/ClassDesign/fileServerModel/README.md", false)
	if err != nil {
		println(err.Error())
		return
	}
	println(file.GetFileData())
}
