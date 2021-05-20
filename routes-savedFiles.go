package main

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
)

func SavedFilesRouter(r *gin.Engine) {
	var listFilenames []string
	imagesLocationDir := "D:\\docs\\films"
	r.GET("/savedfiles", func(c *gin.Context) {
		listFilenames = nil
		files, err := ioutil.ReadDir(imagesLocationDir)
		if err != nil {
			log.Fatal(err)
		}
		for _, f := range files {
			if filepath.Ext(f.Name()) == ".mp4" {
				listFilenames = append(listFilenames, f.Name())
			}
		}
		c.HTML(http.StatusOK, "savedFiles.tmpl", gin.H{
			"listFilenames": listFilenames,
		})
	})
}
