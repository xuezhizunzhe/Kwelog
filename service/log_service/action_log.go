package log_service

import (
	"github.com/gin-gonic/gin"
	"kweblog/models"
	"kweblog/models/enum"
	"net/http"
)

type ActionLog struct {
	c                  *gin.Context
	level              enum.LogLevelType
	title              string
	requestBody        []byte           // 请求体
	responseBody       []byte           // 响应体
	log                *models.LogModel // 用于备份的日志  再创一份日志
	showRequestHeader  bool
	showRequest        bool
	showResponse       bool
	showResponseHeader bool
	itemList           []string    // 日志内容
	responseHeader     http.Header // 响应头
	isMiddlerware      bool        // 是否已经在响应中间件里面调用过save()方法
}

// NewActionLogByGin 在视图层里面获取
func NewActionLogByGin(c *gin.Context) *ActionLog {
	return &ActionLog{
		c: c,
	}

}

func (ac *ActionLog) SetRequest() {

}
