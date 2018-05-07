package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()
	router.GET("/health", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	router.GET("/tasks", getTasks)
	router.POST("/tasks", createTask)
	router.GET("/tasks/:id", getTask)
	router.PUT("/tasks/:id", putTask)
	router.DELETE("/tasks/:id", deleteTask)

	router.Run(":8080")
}

func getTasks(c *gin.Context) {

}

func getTask(c *gin.Context) {

}

func createTask(c *gin.Context) {

}

func putTask(c *gin.Context) {

}

func deleteTask(c *gin.Context) {

}
