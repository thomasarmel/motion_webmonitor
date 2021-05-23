package configread

import (
	"log"
	"net/url"
	"os"
	"path"
	"path/filepath"
)

var ImagesVideosDir string
var ImagesVideosAuthorizedExtensions []string
var CamerasURLs []string
var ViewsDir, PublicDir string
var CommandsStartStopMotion [5]string
var TLSMode bool
var ServerDomain string

func init() {
	ImagesVideosAuthorizedExtensions = append(ImagesVideosAuthorizedExtensions, ".mp4", ".mkv")
	ImagesVideosDir = "C:\\Users\\thoma\\Desktop\\videos"
	CamerasURLs = append(CamerasURLs, "http://192.168.1.20:1941/", "http://192.168.1.25:1941/")
	executableDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	ViewsDir = path.Join(executableDir, "views/")
	PublicDir = path.Join(executableDir, "public/")
	{
		CommandsStartStopMotion[0] = "service motion start"
		CommandsStartStopMotion[1] = "service motion stop"
		CommandsStartStopMotion[2] = "systemctl check motion"
		CommandsStartStopMotion[3] = "active"
		CommandsStartStopMotion[4] = "inactive"
	}
	TLSMode = false
	ServerDomain = "www.example.com"
	checkConfig()
}

func ParseConfigFile(configFile string) {
	//
}

func checkConfig() {
	if _, err := os.Stat(ImagesVideosDir); err != nil {
		log.Fatal("Can't access directory " + ImagesVideosDir)
	}
	for _, cameraURL := range CamerasURLs {
		if _, err := url.ParseRequestURI(cameraURL); err != nil {
			log.Fatal(cameraURL + " is not a valid URL.")
		}
	}
	for _, requiredViewFile := range requiredViewFiles {
		if _, err := os.Stat(path.Join(ViewsDir, requiredViewFile)); err != nil {
			log.Fatal("Can't access file " + requiredViewFile)
		}
	}
}
