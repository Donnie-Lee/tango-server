package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"log"
	"time"
)

type UserClaims struct {
	jwt.StandardClaims //嵌套
	LoginId            int
	UserName           string
}

var JwtKey = []byte("tango20230222")

// GetToken 获取token
func GetToken(loginId int, name string) string {
	//payload
	claims := UserClaims{
		UserName: name,
		LoginId:  loginId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(2 * time.Hour).Unix(), //过期时间
			IssuedAt:  time.Now().Unix(),                    //签发时间
			//Subject:   "user token",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)
	tokenString, err := token.SignedString(JwtKey)
	if err != nil {

		log.Fatal(err.Error())
	}
	return tokenString
}
