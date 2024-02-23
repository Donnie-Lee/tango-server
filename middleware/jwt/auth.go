package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	jwt2 "imserver/common/jwt"
	"imserver/common/response"
	"net/http"
	"strings"
)

var whiteList = [...]string{"/api/account/getCheckCode", "/api/account/loginSms"}

// JwtAuth 验证 token
func JwtAuth(c *gin.Context) {
	for i := range whiteList {
		if strings.HasPrefix(c.Request.RequestURI, whiteList[i]) {
			c.Next()
			return
		}
	}
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		c.JSON(http.StatusOK, response.FailWithMessage("TOKEN不存在"))
		c.Abort() //终止
		return
	}

	claims := jwt2.UserClaims{}
	_, err := jwt.ParseWithClaims(token, &claims, func(t *jwt.Token) (interface{}, error) {
		return jwt2.JwtKey, nil
	})
	//验证失败
	if err != nil {
		ve, _ := err.(*jwt.ValidationError)
		if ve.Errors == jwt.ValidationErrorExpired {
			c.JSON(http.StatusOK, response.FailWithCodeAndMessage(401, "TOKEN已失效"))
		} else {
			c.JSON(http.StatusOK, response.FailWithCodeAndMessage(405, "无效TOKEN"))
		}
		c.Abort() //终止
		return
	}

	c.Set("claims", claims) //传递参数
}
