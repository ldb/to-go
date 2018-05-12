package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"io/ioutil"
	"log"
	"net/http"
)

type server struct {
	Session *mgo.Session
	C       *mgo.Collection
}

func main() {
	gin.SetMode(gin.ReleaseMode)

	s := server{}

	router := gin.Default()

	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	router.GET("/health", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	router.GET("/tasks", s.getTasks)
	router.POST("/tasks", s.createTask)
	router.GET("/tasks/:id", s.getTask)
	router.PUT("/tasks/:id", s.putTask)
	router.DELETE("/tasks/:id", s.deleteTask)

	s.initializeMongoDB()
	router.Run(":8080")
}

func (s *server) initializeMongoDB() {
	type DB struct {
		url        string
		database   string
		collection string
		user       string
		password   string
	}

	db := DB{
		url:        "localhost:27017",
		database:   "to-go",
		collection: "tasks",
		user:       "",
		password:   "",
	}

	session, err := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:    []string{db.url},
		Database: db.database,
		Username: db.user,
		Password: db.password,
	})
	if err != nil {
		log.Fatalf("error connecting to mongoDB %v", err)
	}

	s.Session = session
	s.C = session.DB(db.database).C(db.collection)
}

func (s *server) getTasks(c *gin.Context) {
	err, tasks := s.GetAllTasks()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, tasks)
}

func (s *server) getTask(c *gin.Context) {
	id := c.Param("id")

	err, task := s.FindTask(id)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, task)
}

func (s *server) createTask(c *gin.Context) {
	r := c.Request.Body
	if r != nil {
		c.Status(http.StatusInternalServerError)
	}
	defer r.Close()

	b, err := ioutil.ReadAll(r)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	t := Task{}

	err = json.Unmarshal(b, &t)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	err, task := s.InsertTask(t)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, task)
}

func (s *server) putTask(c *gin.Context) {
	r := c.Request.Body
	if r != nil {
		c.Status(http.StatusInternalServerError)
	}
	defer r.Close()

	b, err := ioutil.ReadAll(r)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	t := Task{}

	err = json.Unmarshal(b, &t)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	id := c.Param("id")
	if id == "" {
		c.AbortWithError(http.StatusBadRequest, nil)
	}

	err, task := s.UpdateTask(id, t)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, task)
}

func (s *server) deleteTask(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.AbortWithError(http.StatusBadRequest, nil)
	}

	err := s.DeleteTask(id)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	c.Status(http.StatusNoContent)
}
