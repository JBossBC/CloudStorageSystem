package util

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

func main() {
	generateToken()
}

func generateToken() {
	cliams := make(jwt.MapClaims)
	cliams["exp"] = time.Now().Unix() + 300000000000000000
	cliams["iat"] = time.Now().Unix()
	cliams["userId"] = "root"
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = cliams
	fmt.Println(token.SignedString([]byte("cloudstoragesystem")))
}
