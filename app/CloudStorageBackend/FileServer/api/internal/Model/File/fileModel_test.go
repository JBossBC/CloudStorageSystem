package File

import (
	"testing"
)

func TestFileFactory_CreateFile(t *testing.T) {
	file, err := GetFileFactory().CreateFile("file:///D:/CloudStorage/ProjectDesign/ClassDesign/fileServerModel/README.md", true)
	if err != nil {
		println(err.Error())
		return
	}
	println(file.GetFileData())
}
