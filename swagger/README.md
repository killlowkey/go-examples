
## Create project
1. swagger-example: store golang code
2. swagger-example/docs: store api data by generated swagger
```shell
mkdir swagger-example
cd swagger-example
go mod init swagger-example
mkdir docs
```

## Install dependencies
1. swag: used to generate api data
2. gin: web framework
3. gin-swagger: used to support gin with swagger
```shell
go install github.com/swaggo/swag/cmd/swag
go get github.com/swaggo/gin-swagger
go get github.com/gin-gonic/gin
```

## Write code
First, need to write `_ "swagger-exmaple/docs/swagger"` that imports swagger dependencies to init.
If you don't init, not access data by generated swagger. 
```golang
package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"net/http"
	_ "swagger-exmaple/docs/swagger"
)

// @title Gin Swagger Example API
// @version 1.0
// @description This is a sample server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:3000
// @BasePath /
// @schemes http
func main() {
	// Gin instance
	r := gin.New()

	// Routes
	r.GET("/", HealthCheck)

	// The url pointing to API definition
	// url := ginSwagger.URL("http://localhost:3000/swagger/doc.json")
	// r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	url := ginSwagger.URL("http://localhost:3000/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	// Start server
	if err := r.Run(":3000"); err != nil {
		log.Fatal(err)
	}
}

// HealthCheck godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router / [get]
func HealthCheck(c *gin.Context) {
	res := map[string]interface{}{
		"data": "Server is up and running",
	}

	c.JSON(http.StatusOK, res)
}
```

## Generate swagger doc
```shell
swag init -g main.go --output docs/swagger
```

## Multiple API feature
Since swag 1.7.9 we are allowing registration of multiple endpoints into the same server.
github.com/swaggo/gin-swagger@v1.6.0

Generate documentation for v1 endpoints
```shell
swag i -g main.go -dir api/v1 --instanceName v1
```


Generate documentation for v2 endpoints
```shell
swag i -g main.go -dir api/v2 --instanceName v2
```

Run example
```shell
    go run main.go
```

## Run project
```shell
go run main.go
```
Access blow link to view swagger website
1. http://localhost:3000/swagger/index.html
2. http://localhost:3000/swagger/doc.json

## Reference
1. [swaggerui-for-gin-go-web-framework](https://levelup.gitconnected.com/tutorial-generate-swagger-specification-and-swaggerui-for-gin-go-web-framework-9f0c038483b5)
2. [Omitting host returns Internal Server Error (500)](https://github.com/swaggo/swag/issues/1019)
3. [declarative-comments-format](https://github.com/swaggo/swag/blob/master/README.md#declarative-comments-format)