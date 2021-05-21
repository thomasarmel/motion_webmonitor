package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"motion_webmonitor/routes"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("views/*")
	r.Static("/images", "public/images")
	r.Static("/script", "public/script")
	r.Static("/style", "public/style")
	r.Static("/fonts", "public/fonts")
	store := cookie.NewStore([]byte("secret")) // TODO: replace by randomly secure string
	r.Use(sessions.Sessions("motion_webmonitor_session", store))
	routes.IndexRoute(r)
	routes.AuthRoute(r)
	routes.DisconnectRoute(r)
	routes.CameraRoute(r)
	routes.SurvRoute(r)
	routes.SavedFilesRouter(r)
	routes.FileViewRoute(r)
	r.Run() // listen and serve on 0.0.0.0:8080
}
