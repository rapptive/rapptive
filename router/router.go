package main

import (
	"net/http"
    "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/rapptive/task"
)

type rapptiveRouter struct {
	port   string
	router gin.Engine
}

func (r *rapptiveRouter) run() {
	r.router.Run(r.port)
}

func main() {
	router := gin.Default()
	// router := rapptiveRouter{port: "8080", router: gin.Default()}

    // Allow http://localhost for development purposes
    config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost"}

    // Tell the router to use the new config
	router.Use(cors.New(config))

	// This handler is responsible for task creation
	router.POST("/tasks", func(c *gin.Context) {
        // Get JSON request and create Task
        var task task.RapptiveTask
        if err := c.ShouldBindJSON(&task); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, task)
	})

	// This handler is responsible for task updates
	router.POST("/tasks/:id", func(c *gin.Context) {
        // Update task
        id := c.Param("id")
        c.JSON(http.StatusOK, gin.H{"status": id})
	})

    // This handler is responsible for reading all tasks
	router.GET("/tasks", func(c *gin.Context) {
        // Get task
        c.JSON(http.StatusOK, "{tasks: 'all tasks'}")
	})

	// This handler is responsible for reading a task
	router.GET("/tasks/:id", func(c *gin.Context) {
        // Get task
        id := c.Param("id")
        c.JSON(http.StatusOK, "{tasks: 'one task " + id + "'}")
	})

	router.Run(":8080")
}
