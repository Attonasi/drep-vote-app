package main

import (
	"example/hello/controller"
	"example/hello/middlewares"
	"example/hello/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Data for populating the template
var (
	voteService    service.VoteService       = service.New()
	voteController controller.VoteController = controller.New(voteService)
)

func main() {
	// Create a new router
	router := gin.New()

	router.Use(gin.Recovery(), middlewares.Logger())

	// Serve static files (like CSS, JavaScript, images)
	router.Static("/static", "./static")

	// Load HTML templates
	router.LoadHTMLGlob("templates/*.html")

	apiRoutes := router.Group("/api")
	{
		// Define your routes
		apiRoutes.GET("/posts", func(c *gin.Context) {
			c.JSON(200, voteController.FindAll(c))
		})

		apiRoutes.POST("/save", func(c *gin.Context) {
			err := voteController.Save(c)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusOK, gin.H{"message": "Video input Valid!"})
			}
		})

		apiRoutes.POST("/build", func(c *gin.Context) {
			err := voteController.Save(c)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusOK, gin.H{"message": "Video input Valid!"})
			}
		})
	}

	viewRoutes := router.Group("/view")
	{
		viewRoutes.GET("/votes", voteController.ShowAll)
	}

	// Start the server
	router.Run(":8080")
}
