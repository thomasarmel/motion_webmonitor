package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func IndexRoute(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", nil)
	})
}
