package routes

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func IndexRoute(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		session := sessions.Default(c)
		if session.Get("connected") == true {
			c.Redirect(http.StatusFound, "/surv")
			return
		}
		c.HTML(http.StatusOK, "index.tmpl", nil)
	})
}
