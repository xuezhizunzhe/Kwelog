package site_api

import (
	"github.com/gin-gonic/gin"
	"kweblog/common/res"
	"kweblog/global"
	"kweblog/middleware"
)

type SiteApi struct {
}

type SiteInfoRequest struct {
	Name string `url:"name"`
}

type SiteInfoResponse struct {
}

func (SiteApi) SiteInfoView(c *gin.Context) {
	var cr SiteInfoRequest
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithError(err, c)
		return
	}
	if cr.Name == "site" {
		global.Config.Site.About.Version = global.Version // 设置版本号
		res.OkWithData(global.Config.Site, c)
		return
	}
	middleware.AdminMiddleware(c)
	_, ok := c.Get("claims")
	if !ok { // 这里的意思是 如果不是管理员的话，那么claims是没有的，所以这里直接返回
		return
	}
	var data any
	// switch 判断值是多少
	switch cr.Name {
	case "email":
		rep := global.Config.Email
		rep.AuthCode = "******" // 将敏感信息隐藏
		data = rep
	case "qq":
		rep := global.Config.QQ
		rep.AppKey = "******" // 将敏感信息隐藏
		data = rep
	case "qiNiu":
		rep := global.Config.Qiniu
		rep.SecretKey = "******" // 将敏感信息隐藏
		data = rep
	case "ai":
		rep := global.Config.Ai
		rep.SecretKey = "******" // 将敏感信息隐藏
		data = rep
	default:
		res.FailWithMsg("不存在的配置", c)
		return
	}
	res.OkWithData(data, c)
}
