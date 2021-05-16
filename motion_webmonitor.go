package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"routes/routes"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("views/*")
	r.Static("/images", "public/images")
	r.Static("/script", "public/script")
	r.Static("/style", "public/style")
	r.Static("/fonts", "public/fonts")
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))
	r.GET("/ping", func(c *gin.Context) {
		session := sessions.Default(c)
		if session.Get("connected") == true {
			c.JSON(200, gin.H{
				"message": "connected !",
			})
		} else {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		}
	})
	/*r.GET("/auth", func(c *gin.Context) {
		session := sessions.Default(c)
		session.Set("connected", true)
		session.Save()
		c.Redirect(302, "/ping")
	})*/
	r.GET("/disconnect", func(c *gin.Context) {
		session := sessions.Default(c)
		session.Clear()
		session.Save()
		c.Redirect(302, "/ping")
	})
	routes.IndexRoute(r)
	routes.AuthRoute(r)
	r.Run() // listen and serve on 0.0.0.0:8080
}
