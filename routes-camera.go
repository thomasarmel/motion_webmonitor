package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func CameraRoute(r *gin.Engine) {

	remote, err := url.Parse("http://192.168.1.20:1941/")
	if err != nil {
		log.Fatal(err)
	}
	proxy := httputil.NewSingleHostReverseProxy(remote)

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

		proxy.Director = func(req *http.Request) {
			req.Header = c.Request.Header
			req.Host = remote.Host
			req.URL.Scheme = remote.Scheme
			req.URL.Host = remote.Host
			req.URL.Path = c.Param("proxyPath")
		}

		proxy.ErrorHandler = func(writer http.ResponseWriter, request *http.Request, err error) {
			c.Header("Cache-Control", "no-cache, private")
			c.Header("Expires", "0")
			c.Header("Max-Age", "0")
			c.Header("Pragma", "no-cache")
			c.File("public/images/cannotconnectmotion.jpg")
			c.Done()
		}
		proxy.ServeHTTP(c.Writer, c.Request)

	})
}
