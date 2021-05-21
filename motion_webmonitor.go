package main

import (
	"crypto/rand"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"log"
	"motion_webmonitor/routes"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	r := gin.Default()
	r.LoadHTMLGlob("views/*")
	r.Static("/images", "public/images")
	r.Static("/script", "public/script")
	r.Static("/style", "public/style")
	r.Static("/fonts", "public/fonts")
	sessionKey, err := generateSessionKey(64)
	if err != nil {
		log.Fatal("Can't generate session key.")
	}
	store := cookie.NewStore(sessionKey)
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

func generateSessionKey(size uint) ([]byte, error) {
	key := make([]byte, size)
	_, err := rand.Read(key)
	return key, err
}
