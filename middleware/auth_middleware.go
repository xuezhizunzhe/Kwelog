package middleware

import (
	"github.com/gin-gonic/gin"
	"kweblog/common/res"
	"kweblog/models/enum"
	"kweblog/utils/jwts"
)

func AdminMiddleware(c *gin.Context) {
	claims, err := jwts.ParseTokenByGin(c) // 解析出来claims的身份
	if err != nil {
		res.FailWithError(err, c)
		c.Abort()
		return
	}
	if claims.Role != enum.AdminRole {
		res.FailWithMsg("权限不足", c)
		c.Abort() // 这个方法用于立即终止当前HTTP请求的处理，并返回给客户端一个响应。
		return
	}
	// TODO 判断这个用户在不在 黑名单里

	c.Set("claims", claims)
}
