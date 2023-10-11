package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
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

func lengthOfLongestSubstring(s string) int {
	subString := make(map[rune]int)
	longest := 0

	for i, v := range s {
		fmt.Println("Loop:", v)
		if j, ok := subString[v]; ok {
			if len(subString) > longest {
				longest = len(subString)
			}
			for k, s := range subString {
				if s <= j {
					delete(subString, k)
				}
			}
			fmt.Println("Emptied sub", subString)
		}
		subString[v] = i
	}

	if len(subString) > longest {
		longest = len(subString)
	}

	return longest
}

func main() {
	fmt.Println(lengthOfLongestSubstring("dvdf"))
}
