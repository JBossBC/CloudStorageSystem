package units

import (
	"crypto/md5"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
)

func Md5(str string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}

// 生成`Token`
func GenerateToken(id int64, account string, name string) (string, error) {
	uc := jwt.MapClaims{ // 声明 payload
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uc) // 创建使用HS256签名的JWT
	signedToken, err := token.SignedString("123")          // 签名,组装token
	return signedToken, err
}
