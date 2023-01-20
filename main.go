package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("templates/*")

	//routers
	//router.GET("/", func(c *gin.Context) {
	//	// Call the HTML method of the Context to render a template
	//	c.HTML(
	//		// Set the HTTP status to 200 (OK)
	//		http.StatusOK,
	//		// Use the index.html template
	//		"index.html",
	//		// Pass the data that the page uses (in this case, 'title')
	//		gin.H{
	//			"title": "Home Page",
	//		},
	//	)
	//
	//})

	router.GET("/", showIndexPage)
	router.GET("/article/view/:article_id", getArticle)
	// Start serving the application
	router.Run()
}
