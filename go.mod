module motion_webmonitor

go 1.15

require (
	github.com/asaskevich/govalidator v0.0.0-20210307081110-f21760c49a8d
	github.com/gabriel-vasile/mimetype v1.4.2
	github.com/gin-contrib/sessions v0.0.5
	github.com/gin-gonic/autotls v0.0.5
	github.com/gin-gonic/gin v1.9.1
	github.com/kr/pretty v0.3.0 // indirect
	github.com/rogpeppe/go-internal v1.8.0 // indirect
	github.com/ugorji/go v1.2.7 // indirect
	golang.org/x/crypto v0.9.0
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
)

replace github.com/gin-gonic/autotls => github.com/thomasarmel/autotls v0.0.0-20210527074749-e77b44254795

//replace github.com/gin-gonic/autotls => ../autotls
