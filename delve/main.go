package main

import "github.com/gin-gonic/gin"

// main
// go install github.com/go-delve/delve/cmd/dlv@latest
// go build -gcflags "all=-N -l"
// dlv --listen=:3000 --headless=true --api-version=2 --accept-multiclient exec ./delve.exe
func main() {
	r := gin.Default()
	r.GET("/ping", ping)
	r.Run(":8080")
}

func ping(ctx *gin.Context) {
	data := make(map[string]any)
	data["code"] = 200
	data["msg"] = "pong"
	ctx.JSON(200, data)
}
