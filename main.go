package main

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"log"
	"net/http"
)

const (
	url        = "localhost:27017"
	database   = "todo"
	collection = "tasks"
	user       = ""
	password   = ""
)

var session *mgo.Session

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

	initializeMongoDB()
	router.Run(":8080")
}

func initializeMongoDB() {
	s, err := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:    []string{url},
		Database: database,
		Username: user,
		Password: password,
	})
	if err != nil {
		log.Fatalf("error connecting to mongoDB %v", err)
	}

	session = s
}

func getTasks(c *gin.Context) {
	err, tasks := GetAllTasks()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, tasks)
}

func getTask(c *gin.Context) {

}

func createTask(c *gin.Context) {

}

func putTask(c *gin.Context) {

}

func deleteTask(c *gin.Context) {

}
