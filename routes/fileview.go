package routes

import (
	"fmt"
	"github.com/gabriel-vasile/mimetype"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func FileViewRoute(r *gin.Engine) {
	r.GET("/fileview", func(c *gin.Context) {
		filename, paramExists := c.GetQuery("file")
		if !paramExists {
			c.String(http.StatusBadRequest, "Please specify a file param.")
			return
		}
		if filepath.Ext(filename) != ".mp4" || filename[0] == '.' {
			c.String(http.StatusUnauthorized, "Incorrect file format.")
			return
		}
		pathToFile := "D:\\docs\\films\\" + filename
		file, fileErr := os.Stat(pathToFile)
		if fileErr != nil {
			c.String(http.StatusNotFound, "Can't access specified file.")
			return
		}
		c.Header("Accept-Ranges", "bytes")
		httpRange := c.GetHeader("Range")
		fileReader, err := os.Open(pathToFile)
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
			fmt.Println(httpRange)
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
		mime, _ := mimetype.DetectFile(pathToFile)
		c.DataFromReader(httpStatus, length, mime.String(), fileReader, map[string]string{
			"Content-Range": "bytes " + strconv.FormatInt(start, 10) + "-" + strconv.FormatInt(end, 10) + "/" + strconv.FormatInt(size, 10),
		})

	})
}
