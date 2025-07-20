package router

import (
	"github.com/gin-gonic/gin"
	"kweblog/api"
)

func SiteRouter(r *gin.RouterGroup) {
	app := api.App.SiteApi
	r.GET("/site/:name", app.SiteInfoView)
}
