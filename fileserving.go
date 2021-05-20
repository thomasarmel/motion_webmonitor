package main

import (
	"github.com/gabriel-vasile/mimetype"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func ServeVideo(c *gin.Context, videoPath string) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("panic occurred:", err)
		}
	}()
	file, fileErr := os.Stat(videoPath)
	if fileErr != nil {
		c.String(http.StatusNotFound, "Can't access specified file.")
		return
	}
	c.Header("Accept-Ranges", "bytes")
	httpRange := c.GetHeader("Range")
	fileReader, err := os.Open(videoPath)
	var (
		size   int64 = file.Size()
		length int64 = size
		start  int64 = 0
		end    int64 = size - 1
	)
	if err != nil {
		c.String(http.StatusNotAcceptable, "Can't open specified file.")
		return
	}
	httpStatus := http.StatusOK
	if httpRange != "" {
		rangestr := strings.Split(httpRange, "=")
		if len(rangestr) != 2 {
			c.Header("Content-Range", "bytes "+strconv.FormatInt(start, 10)+"-"+strconv.FormatInt(end, 10)+"/"+strconv.FormatInt(size, 10))
			c.String(http.StatusRequestedRangeNotSatisfiable, "Can't access specified range.")
			return
		}
		rangeArray := strings.Split(rangestr[1], "-")
		if len(rangeArray) == 0 {
			c.Header("Content-Range", "bytes "+strconv.FormatInt(start, 10)+"-"+strconv.FormatInt(end, 10)+"/"+strconv.FormatInt(size, 10))
			c.String(http.StatusRequestedRangeNotSatisfiable, "Can't access specified range.")
			return
		}
		cstart, err := strconv.ParseInt(rangeArray[0], 10, 64)
		if err != nil {
			c.Header("Content-Range", "bytes "+strconv.FormatInt(start, 10)+"-"+strconv.FormatInt(end, 10)+"/"+strconv.FormatInt(size, 10))
			c.String(http.StatusRequestedRangeNotSatisfiable, "Can't access specified range.")
			return
		}
		cend := end
		if len(rangeArray) > 1 && rangeArray[1] != "" {
			cend, err = strconv.ParseInt(rangeArray[1], 10, 64)
			if err != nil {
				c.Header("Content-Range", "bytes "+strconv.FormatInt(start, 10)+"-"+strconv.FormatInt(end, 10)+"/"+strconv.FormatInt(size, 10))
				c.String(http.StatusRequestedRangeNotSatisfiable, "Can't access specified range.")
				return
			}
		}
		if cstart > cend || cstart > size-1 || cend >= size {
			c.Header("Content-Range", "bytes "+strconv.FormatInt(start, 10)+"-"+strconv.FormatInt(end, 10)+"/"+strconv.FormatInt(size, 10))
			c.String(http.StatusRequestedRangeNotSatisfiable, "Can't access specified range.")
			return
		}

		c.Status(http.StatusPartialContent)
		start = cstart
		end = cend
		length = end - start + 1
		fileReader.Seek(start, 0)
		httpStatus = http.StatusPartialContent
	}
	mime, _ := mimetype.DetectFile(videoPath)
	c.DataFromReader(httpStatus, length, mime.String(), fileReader, map[string]string{
		"Content-Range": "bytes " + strconv.FormatInt(start, 10) + "-" + strconv.FormatInt(end, 10) + "/" + strconv.FormatInt(size, 10),
	})
}
