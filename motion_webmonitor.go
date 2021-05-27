package main

import (
	"crypto/rand"
	"crypto/tls"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/acme/autocert"
	"log"
	"motion_webmonitor/configread"
	"motion_webmonitor/routes"
	"net/http"
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
	gin.SetMode(gin.ReleaseMode)
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
	store.Options(sessions.Options{
		MaxAge:   3600,
		Secure:   configread.TLSMode,
		HttpOnly: true,
		Path:     "/",
		SameSite: http.SameSiteStrictMode,
	})
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
	if configread.TLSMode {
		cfg := tls.Config{
			MinVersion:               tls.VersionTLS12,
			CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
			PreferServerCipherSuites: true,
			CipherSuites: []uint16{
				tls.TLS_CHACHA20_POLY1305_SHA256,
				tls.TLS_AES_128_GCM_SHA256,
				tls.TLS_AES_256_GCM_SHA384,
				tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
				tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
				//tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
				tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
				//tls.TLS_RSA_WITH_AES_256_CBC_SHA,
			},
		}
		m := autocert.Manager{
			Prompt:     autocert.AcceptTOS,
			HostPolicy: autocert.HostWhitelist(configread.ServerDomains...),
		}
		log.Fatal(autotls.RunWithManagerAndTLSConfig(r, &m, cfg))

	} else {
		r.Run(":" + strconv.Itoa(int(configread.NotSecureModePort)))
	}
}

func generateSessionKey(size uint) ([]byte, error) {
	key := make([]byte, size)
	_, err := rand.Read(key)
	return key, err
}
