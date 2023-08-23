package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"net/http"
	_ "swagger-exmaple/docs/swagger"
	"swagger-exmaple/dto/req"
	"swagger-exmaple/dto/resp"
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
	r.GET("/users", GetUserList)

	// change swagger info
	// swagger.SwaggerInfo

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

// GetUserList godoc
// @Summary get all user
// @Description get all user
// @Tags user
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param page query int true  "page"
// @Param size query int true  "size"
// @Success 200 {array} resp.UserResp
// @failure 400 {object} resp.CommandResp
// @failure 400 {object} resp.CommandResp
// @Router /users [get]
func GetUserList(c *gin.Context) {
	var userReq req.UserListReq
	if err := c.ShouldBindQuery(&userReq); err != nil {
		c.JSON(http.StatusBadRequest, resp.CommandResp{
			Code: 400,
			Msg:  "bad request params",
		})
		return
	}

	c.JSON(http.StatusOK, []resp.UserResp{
		{
			Name: "ray",
			Age:  10,
		},
		{
			Name: "run",
			Age:  98,
		},
	})
}
