package routes

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"motion_webmonitor/configread"
	"net/http"
	"path/filepath"
)

func SavedFilesRouter(r *gin.Engine) {
	var listFilenames []string
	r.GET("/savedfiles", func(c *gin.Context) {
		session := sessions.Default(c)
		if session.Get("connected") != true {
			c.Redirect(http.StatusFound, "/?e=2")
			return
		}
		listFilenames = nil
		files, err := ioutil.ReadDir(configread.ImagesVideosDir)
		if err != nil {
			log.Fatal(err)
		}
		for _, f := range files {
			if configread.Contains(configread.ImagesVideosAuthorizedExtensions, filepath.Ext(f.Name())) {
				listFilenames = append(listFilenames, f.Name())
			}
		}
		c.HTML(http.StatusOK, "savedFiles.tmpl", gin.H{
			"listFilenames": listFilenames,
		})
	})
}
