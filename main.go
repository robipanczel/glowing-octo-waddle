package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "pong")
	})

	anime := r.Group("/hello")
	{
		anime.GET("", func(ctx *gin.Context) {
			ctx.IndentedJSON(http.StatusOK, "HELLO WORLD")
		})
		anime.POST("", func(ctx *gin.Context) {
			ctx.IndentedJSON(http.StatusOK, "HELLO WORLD")
		})
		anime.DELETE("", func(ctx *gin.Context) {
			ctx.IndentedJSON(http.StatusOK, "HELLO WORLD")
		})
	}

	return r
}

func main() {
	r := setupRouter()

	r.Run(":8080")
}
