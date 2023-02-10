module motion_webmonitor

go 1.15

require (
	github.com/asaskevich/govalidator v0.0.0-20210307081110-f21760c49a8d
	github.com/gabriel-vasile/mimetype v1.4.1
	github.com/gin-contrib/sessions v0.0.5
	github.com/gin-gonic/autotls v0.0.5
	github.com/gin-gonic/gin v1.8.2
	github.com/go-playground/validator/v10 v10.11.2 // indirect
	github.com/goccy/go-json v0.10.0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/mattn/go-isatty v0.0.17 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/ugorji/go/codec v1.2.9 // indirect
	golang.org/x/crypto v0.6.0
	gopkg.in/go-playground/assert.v1 v1.2.1 // indirect
	gopkg.in/go-playground/validator.v9 v9.29.1 // indirect
)

replace github.com/gin-gonic/autotls => github.com/thomasarmel/autotls v0.0.0-20210527074749-e77b44254795

//replace github.com/gin-gonic/autotls => ../autotls
