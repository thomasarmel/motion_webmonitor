package configread

var ImagesVideosDir string
var ImagesVideosAuthorizedExtensions []string
var CamerasURLs []string

func init() {
	ImagesVideosAuthorizedExtensions = append(ImagesVideosAuthorizedExtensions, ".mp4", ".mkv")
	ImagesVideosDir = "D:\\docs\\films"
	CamerasURLs = append(CamerasURLs, "http://192.168.1.20:1941/", "http://192.168.1.25:1941/")
}

func ParseConfigFile(configFile string) {
	//
}
