package routes

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"motion_webmonitor/configread"
	"net/http"
	"os/exec"
	"strings"
)

func StartStopMotionRoute(r *gin.Engine) {
	r.GET("/startstopmotion", func(c *gin.Context) {
		session := sessions.Default(c)
		if session.Get("connected") != true {
			c.Redirect(http.StatusFound, "/?e=2")
			return
		}
		action, hasAction := c.GetQuery("action")
		if !hasAction {
			c.String(http.StatusBadRequest, "Please specify an action param.")
			return
		}
		switch action {
		case "check":
			checkMotionStatus(c)
			break
		case "start":
			startStopMotion(true, c)
			break
		case "stop":
			startStopMotion(false, c)
			break
		default:
			c.String(http.StatusBadRequest, "Unrecognized action param.")
			break
		}
	})
}

func startStopMotion(start bool, c *gin.Context) {
	if configread.CommandsStartStopMotion[3] == "" || configread.CommandsStartStopMotion[4] == "" {
		c.String(http.StatusNotImplemented, "Not implemented.")
		return
	}
	var cmdSplitted []string
	if start {
		cmdSplitted = strings.Split(configread.CommandsStartStopMotion[3], " ")
	} else {
		cmdSplitted = strings.Split(configread.CommandsStartStopMotion[4], " ")
	}
	cmd := exec.Command(cmdSplitted[0], cmdSplitted[1:]...)
	_, err := cmd.CombinedOutput()
	if err != nil {
		if _, ok := err.(*exec.ExitError); ok {
			c.String(http.StatusExpectationFailed, "inactive")
			return
		} else {
			c.String(http.StatusFailedDependency, "Failed to run start/stop command.")
			return
		}
	}
	c.String(http.StatusOK, "Done")
}

func checkMotionStatus(c *gin.Context) {
	if configread.CommandsStartStopMotion[2] == "" {
		c.String(http.StatusNotImplemented, "unknown")
		return
	}
	cmdSplitted := strings.Split(configread.CommandsStartStopMotion[2], " ")
	cmd := exec.Command(cmdSplitted[0], cmdSplitted[1:]...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		if _, ok := err.(*exec.ExitError); ok {
			c.String(http.StatusOK, "inactive")
			return
		} else {
			c.String(http.StatusFailedDependency, "Failed to run check command.")
			return
		}
	}
	switch strings.TrimSpace(string(out)) {
	case configread.CommandsStartStopMotion[3]:
		c.String(http.StatusOK, "active")
		break
	case configread.CommandsStartStopMotion[4]:
		c.String(http.StatusOK, "inactive")
		break
	default:
		c.String(http.StatusOK, "unknown")
		break
	}
}
