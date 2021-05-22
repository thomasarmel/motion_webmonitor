package routes

import (
	"bufio"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"os"
	"strings"
)

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func AuthRoute(r *gin.Engine) {
	file, err := os.Open(".passwd")
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)
	scanner := bufio.NewScanner(file)
	users := make(map[string]string)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) != "" {
			arr := strings.Split(line, ":")
			if len(arr) != 2 {
				log.Fatal("Error: wrong .passwd file format")
			}
			users[arr[0]] = arr[1]
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	r.POST("/auth", func(c *gin.Context) {
		loginHTTP, ok := c.GetPostForm("login")
		if !ok {
			c.Redirect(http.StatusFound, "/?e=1")
			return
		}
		passwordHTTP, ok := c.GetPostForm("password")
		if !ok {
			c.Redirect(http.StatusFound, "/?e=1")
			return
		}
		pass, ok := users[loginHTTP]
		if !ok {
			c.Redirect(http.StatusFound, "/?e=2")
			return
		}
		if !checkPasswordHash(passwordHTTP, pass) {
			c.Redirect(http.StatusFound, "/?e=2")
			return
		}
		session := sessions.Default(c)
		session.Set("connected", true)
		err := session.Save()
		if err != nil {
			c.Redirect(http.StatusFound, "/?e=3")
			return
		}
		c.Redirect(http.StatusFound, "/surv")
	})
}
