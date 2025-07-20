package res

import (
	"github.com/gin-gonic/gin"
	"kweblog/utils/validate"
)

type Code int

const (
	SuccessCode     Code = 0
	FailValidCode   Code = 1001
	FailServiceCode Code = 1002
)

func (c Code) String() string {
	switch c {
	case SuccessCode:
		return "成功"
	case FailValidCode:
		return "校验失败" // 校验失败
	case FailServiceCode:
		return "服务异常" // 服务异常
	}
	return ""
}

var empty = map[string]any{}

type Response struct {
	Code Code   `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

func (r Response) Json(c *gin.Context) {
	c.JSON(200, r)
}

func Ok(data any, msg string, c *gin.Context) {
	Response{
		Code: SuccessCode,
		Data: data,
		Msg:  msg,
	}.Json(c)
}

func OkWithData(data any, c *gin.Context) {
	Response{SuccessCode, data, "成功"}.Json(c)

}

func OkWithList(list any, count int, c *gin.Context) { // 总页  count 是总条数
	Response{SuccessCode, map[string]any{
		"list":  list,  // 列表数据
		"count": count, // 总条数
	}, "成功"}.Json(c)
}

func OkWithMsg(msg string, c *gin.Context) {
	Response{SuccessCode, empty, msg}.Json(c)

}

func FailWithMsg(msg string, c *gin.Context) {
	Response{FailValidCode, empty, msg}.Json(c) // 校验
}

func FailWithData(data any, msg string, c *gin.Context) {
	Response{FailServiceCode, data, msg}.Json(c) // 服务
}

func FailWithCode(code Code, c *gin.Context) { // 这里的msg需要用code.String()来获取  通过code来获取错误信息
	Response{SuccessCode, empty, code.String()}.Json(c)
}

func FailWithError(err error, c *gin.Context) {
	data, msg := validate.ValidateError(err) // 错误信息显示中文
	FailWithData(data, msg, c)
}
