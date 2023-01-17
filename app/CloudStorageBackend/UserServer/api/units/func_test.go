package units

import (
	"fmt"
	"testing"
)

func TestMd5(t *testing.T) {
	cp := Md5("123456")
	fmt.Println(cp)
}
