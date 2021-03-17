package main

import (
	// "fmt"

	"github.com/gin-gonic/gin"
	"github.com/shubhamsnehi/gin-demo/controllers"
	"github.com/shubhamsnehi/gin-demo/service"
)

var (
	userService    service.UserService        = service.New()
	UserController controllers.UserController = controllers.New(userService)
)

func main() {
	r := gin.Default()
	r.GET("/owner", func(ctx *gin.Context) {
		ctx.JSON(200, UserController.ShowOwners())
	})
	r.GET("/ownerid", func(ctx *gin.Context) {
		ctx.JSON(200, UserController.QueryOwners1(ctx))
	})
	r.POST("/owner", func(ctx *gin.Context) {
		ctx.JSON(200, UserController.PostOwner(ctx))
	})

	//----------------------------------------------
	r.GET("/users", func(ctx *gin.Context) {
		ctx.JSON(200, UserController.ShowUsers())
	})
	r.POST("/users", func(ctx *gin.Context) {
		ctx.JSON(200, UserController.Add(ctx))
	})
	r.PATCH("/users", func(ctx *gin.Context) {
		ctx.JSON(200, UserController.Update(ctx))
	})
	r.DELETE("/users", func(ctx *gin.Context) {
		ctx.JSON(200, UserController.Delete(ctx))
	})
	r.GET("/usr", controllers.QueryString())

	r.GET("/param/:fname/:lname", func(ctx *gin.Context) {
		ctx.JSON(200, UserController.ParamString(ctx))
	})
	r.Run()
}

// r.GET("/user/:name", func(c *gin.Context) {
// 	name := c.Param("name")
// 	c.String(http.StatusOK, "Hello %s", name)
// })
// r.GET("/welcome", func(c *gin.Context) {
// 	firstname := c.DefaultQuery("firstname", "Guest")
// 	lastname := c.Query("lastname") // shortcut for c.Request.URL.Query().Get("lastname")

// 	c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
// })
