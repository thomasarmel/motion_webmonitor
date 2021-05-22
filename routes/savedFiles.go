package routes

import (
	"crypto/rand"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"math/big"
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
			c.String(http.StatusPreconditionFailed, "Can't open images directory")
			return
		}
		for _, f := range files {
			if configread.Contains(configread.ImagesVideosAuthorizedExtensions, filepath.Ext(f.Name())) {
				listFilenames = append(listFilenames, f.Name())
			}
		}
		cleanFilesToken, e := generateRandomString(32)
		if e != nil {
			c.String(http.StatusExpectationFailed, "Can't generate clean files token.")
			return
		}
		session.Set("cleanfilestoken", cleanFilesToken)
		err = session.Save()
		if err != nil {
			c.String(http.StatusExpectationFailed, "Can't save clean files token on session.")
			return
		}
		c.HTML(http.StatusOK, "savedFiles.tmpl", gin.H{
			"listFilenames":   listFilenames,
			"cleanFilesToken": cleanFilesToken,
		})
	})
}

func generateRandomString(n int) (string, error) {
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-"
	ret := make([]byte, n)
	for i := 0; i < n; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return "", err
		}
		ret[i] = letters[num.Int64()]
	}

	return string(ret), nil
}
