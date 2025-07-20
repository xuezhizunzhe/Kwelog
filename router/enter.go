package router

import (
	"github.com/gin-gonic/gin"
	"kweblog/global"
)

func Run() {
	gin.SetMode(global.Config.System.GinMode)

	r := gin.Default()
	nr := r.Group("/api")

	r.Run(global.Config.System.Ip + ":" + global.Config.System.Port)
}
