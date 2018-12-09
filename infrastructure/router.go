package infrastructure

import (
	"github.com/gin-gonic/gin"
	"github.com/yoshikawataiki/simple-api/interfaces/controller"
)

// Router is gin engine
var Router *gin.Engine

func init() {
	router := gin.Default()

	userController := controller.NewUserController(NewSQLHandler())

	// user api route
	router.POST("/users", func(c *gin.Context) { userController.Create(c) })
	router.GET("/users", func(c *gin.Context) { userController.Index(c) })
	router.GET("/users/:id", func(c *gin.Context) { userController.Show(c) })

	Router = router
}
