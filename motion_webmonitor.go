package main

import (
	"crypto/rand"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"log"
	"motion_webmonitor/configread"
	"motion_webmonitor/routes"
	"os"
	"path"
	"runtime"
	"strconv"
)

func main() {
	if len(os.Args) > 1 {
		configread.ParseConfigFile(os.Args[1])
	} else {
		configread.CheckConfig()
	}
	runtime.GOMAXPROCS(runtime.NumCPU())
	r := gin.Default()
	r.LoadHTMLGlob(path.Join(configread.ViewsDir, "*"))
	r.Static("/images", path.Join(configread.PublicDir, "images"))
	r.Static("/script", path.Join(configread.PublicDir, "script"))
	r.Static("/style", path.Join(configread.PublicDir, "style"))
	r.Static("/fonts", path.Join(configread.PublicDir, "fonts"))
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
	routes.CleanFilesRouter(r)
	routes.StartStopMotionRoute(r)
	r.Run(":" + strconv.Itoa(int(configread.NotSecureModePort)))
}

func generateSessionKey(size uint) ([]byte, error) {
	key := make([]byte, size)
	_, err := rand.Read(key)
	return key, err
}
