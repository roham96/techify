package main

import "github.com/gin-gonic/gin"


func main() {
	router := gin.Default()
	router.GET("/articles", getArticles)
	router.GET("/articles/:id", getArticleByID)
	router.POST("/articles", postArticles)

	router.Run("localhost:8000")
}

