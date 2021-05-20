package routes

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SurvRoute(r *gin.Engine) {
	r.GET("/surv", func(c *gin.Context) {
		session := sessions.Default(c)
		if session.Get("connected") != true { // TODO: everywhere
			c.Redirect(http.StatusFound, "/")
			return
		}
		c.HTML(http.StatusOK, "surv.tmpl", nil)
	})
}
