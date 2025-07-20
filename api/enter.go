package api

import "kweblog/api/site_api"

type Api struct {
	SiteApi site_api.SiteApi
}

var App = Api{}
