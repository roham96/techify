package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

// getArticles responds with the list of all articles as JSON.
func getArticles(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, articles)
}

// postArticles adds an article from JSON received in the request body.
func postArticles(c *gin.Context) {
	var newArticle Article

	// Call BindJSON to bind the received JSON to
	// newArticle.
	if err := c.BindJSON(&newArticle); err != nil {
		return
	}

	// Add the new article to the slice.
	articles = append(articles, newArticle)
	c.IndentedJSON(http.StatusCreated, newArticle)
}

// getArticleByID locates the article whose ID value matches the id
// parameter sent by the client, then returns that article as a response.
func getArticleByID(c *gin.Context) {
	id := c.Param("id")

	// Loop through the list of articles, looking for
	// an article whose ID value matches the parameter.
	for _, a := range articles {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "article not found"})
}