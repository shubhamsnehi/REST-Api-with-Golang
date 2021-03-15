package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/shubhamsnehi/gin-demo/entity"
	"github.com/shubhamsnehi/gin-demo/service"
)

type UserController interface {
	ShowUsers() []entity.User
	Add(c *gin.Context) entity.User
	Update(c *gin.Context) entity.User
	Delete(c *gin.Context) entity.User
	// QueryString(c *gin.Context) entity.User
	ParamString(c *gin.Context) entity.User
}

type controller struct {
	service service.UserService
}

func New(service service.UserService) UserController {
	return controller{
		service: service,
	}
}

func (c controller) ShowUsers() []entity.User {
	return c.service.ShowUsers()
}

func (c controller) Add(ctx *gin.Context) entity.User {
	var user entity.User
	ctx.BindJSON(&user)
	c.service.Add(user)
	return user
}

func (c controller) Update(ctx *gin.Context) entity.User {
	var user entity.User
	ctx.BindJSON(&user)
	c.service.Update(user)
	return user
}

func (c controller) Delete(ctx *gin.Context) entity.User {
	var user entity.User
	ctx.BindJSON(&user)
	c.service.Delete(user)
	return user
}

func QueryString() gin.HandlerFunc {
	return func(ctx *gin.Context){
	var user entity.User
	user.ID = ctx.Query("id")
	user.Name = ctx.Query("name")
	// c.service.Add(user)
	ctx.JSON(200, user)
	}
}

func (c controller) ParamString(ctx *gin.Context) entity.User {
	var user entity.User
	user.ID = ctx.Param("id")
	user.Name = ctx.Param("name")
	c.service.Add(user)
	return user
}
