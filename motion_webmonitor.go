package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))
	r.GET("/ping", func(c *gin.Context) {
		session := sessions.Default(c)
		if session.Get("connected") == "true" {
			c.JSON(200, gin.H{
				"message": "connected !",
			})
		} else {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		}
	})
	r.GET("/auth", func(c *gin.Context) {
		session := sessions.Default(c)
		session.Set("connected", "true")
		c.Redirect(302, "/ping")
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
