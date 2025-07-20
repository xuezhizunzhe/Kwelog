package jwts

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"kweblog/global"
	"kweblog/models/enum"
	"strings"
)

type Claims struct {
	UserID   uint          `json:"userID"`
	UserName string        `json:"userName"`
	Role     enum.RoleType `json:"role"` // 用户角色，自定义的enum.RoleType类型
}

// MyClaims 这样，MyClaims 既包含了自定义的用户信息，也包含了 JWT 的标准字段
type MyClaims struct {
	Claims             // 嵌入自定义声明
	jwt.StandardClaims // 嵌入JWT的标准声明
}

func ParseTokenByGin(c *gin.Context) (*MyClaims, error) {
	token := c.GetHeader("token")
	if token == "" {
		token = c.Query("token")
	}
	return ParseToken(token)
}

func ParseToken(tokenString string) (*MyClaims, error) {
	if tokenString == "" {
		return nil, errors.New("请登录")
	}
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(global.Config.Jwt.Secret), nil
	})
	if err != nil {
		if strings.Contains(err.Error(), "token is expired") { // 判断 这个err里面有没有包含对应 "token is expired" 的子字符串
			return nil, errors.New("token 过期")
		}
		if strings.Contains(err.Error(), "signature is invalid") {
			return nil, errors.New("token 无效")
		}
		if strings.Contains(err.Error(), "token contains is invalid") {
			return nil, errors.New("token 非法")
		}
		fmt.Println(err, 1)
		return nil, err
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("token 无效")

}
