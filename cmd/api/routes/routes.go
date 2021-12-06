package router

import (
	"zzucturum-app/cmd/api/middleware"
	"zzucturum-app/pkg/http-transport/router"
)

var Routes = []router.Route{
	router.NewRoute("POST", "/socaial/media/add", middleware.AddSocialMedia),
	router.NewRoute("POST", "/social/media/edit", middleware.EditSocialMedia),
	router.NewRoute("GET", "/social/media/get", middleware.GetSocialMedia),
}
