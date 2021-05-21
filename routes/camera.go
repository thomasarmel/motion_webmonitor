package routes

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
	"motion_webmonitor/configread"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func CameraRoute(r *gin.Engine) {
	var remotes []*url.URL
	var proxys []*httputil.ReverseProxy
	for _, u := range configread.CamerasURLs {
		r, err := url.Parse(u)
		if err != nil {
			log.Fatal(err)
		}
		remotes = append(remotes, r)
		proxys = append(proxys, httputil.NewSingleHostReverseProxy(r))
	}

	r.GET("/camera/:id", func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Println("panic occurred:", err)
			}
		}()

		session := sessions.Default(c)
		if session.Get("connected") != true {
			c.Redirect(http.StatusFound, "/?e=2")
			return
		}

		var cam camera
		if err := c.ShouldBindUri(&cam); err != nil {
			c.String(http.StatusBadRequest, "Error camera number format.")
			return
		}
		if cam.ID < 0 || cam.ID >= len(proxys) {
			c.String(http.StatusNotAcceptable, "Can't access specified camera.")
			return
		}

		proxys[cam.ID].Director = func(req *http.Request) {
			req.Header = c.Request.Header
			req.Host = remotes[cam.ID].Host
			req.URL.Scheme = remotes[cam.ID].Scheme
			req.URL.Host = remotes[cam.ID].Host
			req.URL.Path = c.Param("proxyPath")
		}

		proxys[cam.ID].ErrorHandler = func(writer http.ResponseWriter, request *http.Request, err error) {
			c.Header("Cache-Control", "no-cache, private")
			c.Header("Expires", "0")
			c.Header("Max-Age", "0")
			c.Header("Pragma", "no-cache")
			c.File("public/images/cannotconnectmotion.jpg")
			c.Done()
		}
		proxys[cam.ID].ServeHTTP(c.Writer, c.Request)

	})
}

type camera struct {
	ID int `uri:"id"`
}
