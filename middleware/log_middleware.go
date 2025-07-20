package middleware

import (
	"github.com/gin-gonic/gin"
	"kweblog/service/log_service"
	"net/http"
)

type ResponseWriter struct {
	gin.ResponseWriter
	Body []byte
	Head http.Header
}

func (w *ResponseWriter) Writer(data []byte) (int, error) {
	w.Body = append(w.Body, data...)
	return w.ResponseWriter.Write(data)
}

func (w *ResponseWriter) Header() http.Header {
	return w.Head
}
func LogMiddleware(c *gin.Context) {
	log := log_service.NewActionLogByGin(c) // 在请求中间件中创建日志对象

}
