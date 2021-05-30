package routes

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"motion_webmonitor/configread"
	"net/http"
)

func SurvRoute(r *gin.Engine) {
	var numCams []int = nil
	if len(configread.CamerasURLs) > 0 {
		numCams = makeRange(0, len(configread.CamerasURLs)-1)
	}
	hasSavesDir := (configread.ImagesVideosDir != "")

	r.GET("/surv", func(c *gin.Context) {
		session := sessions.Default(c)
		if session.Get("connected") != true {
			c.Redirect(http.StatusFound, "/?e=2")
			return
		}
		startStopMotionToken, e := generateRandomString(32)
		if e != nil {
			c.String(http.StatusExpectationFailed, "Can't generate clean files token.")
			return
		}

		session.Set("startstopmotiontoken", startStopMotionToken)
		err := session.Save()
		if err != nil {
			c.String(http.StatusExpectationFailed, "Can't save clean files token on session.")
			return
		}
		c.HTML(http.StatusOK, "surv.tmpl", gin.H{
			"numCams":              numCams,
			"startstopmotiontoken": startStopMotionToken,
			"hasSavesDir":          hasSavesDir,
		})
	})
}

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}
