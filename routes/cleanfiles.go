package routes

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"motion_webmonitor/configread"
	"net/http"
	"os"
	"path"
	"path/filepath"
)

func CleanFilesRouter(r *gin.Engine) {
	r.GET("/cleanfiles", func(c *gin.Context) {
		session := sessions.Default(c)
		if session.Get("connected") != true {
			c.Redirect(http.StatusFound, "/?e=2")
			return
		}
		token, hasToken := c.GetQuery("token")
		if !hasToken {
			c.String(http.StatusBadRequest, "Please specify a token for this request.")
			return
		}
		if token != session.Get("cleanfilestoken") {
			c.String(http.StatusForbidden, "Bad security token.")
			return
		}
		files, err := ioutil.ReadDir(configread.ImagesVideosDir)
		if err != nil {
			c.String(http.StatusPreconditionFailed, "Can't open images directory")
			return
		}
		var notRemovedFiles uint32 = 0
		for _, f := range files {
			if configread.Contains(configread.ImagesVideosAuthorizedExtensions, filepath.Ext(f.Name())) {
				if err := os.Remove(path.Join(configread.ImagesVideosDir, f.Name())); err != nil {
					notRemovedFiles++
				}
			}
		}
		if notRemovedFiles != 0 {
			c.String(http.StatusExpectationFailed, "Not all files could be deleted.")
			return
		}
		c.String(http.StatusOK, "Done, all files deleted.")
	})
}
