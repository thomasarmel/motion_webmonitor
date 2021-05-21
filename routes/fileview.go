package routes

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
	"motion_webmonitor/configread"
	"motion_webmonitor/fileserving"
	"net/http"
	"path"
	"path/filepath"
)

func FileViewRoute(r *gin.Engine) {
	r.GET("/fileview", func(c *gin.Context) {
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
		filename, paramExists := c.GetQuery("file")
		if !paramExists {
			c.String(http.StatusBadRequest, "Please specify a file param.")
			return
		}
		if !configread.Contains(configread.ImagesVideosAuthorizedExtensions, filepath.Ext(filename)) || filename[0] == '.' {
			c.String(http.StatusUnauthorized, "Incorrect file format.")
			return
		}
		pathToFile := path.Join(configread.ImagesVideosDir, filename)
		fileserving.ServeVideo(c, pathToFile)
	})
}
