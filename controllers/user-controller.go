package controllers

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/shubhamsnehi/gin-demo/entity"
	"github.com/shubhamsnehi/gin-demo/service"
)

type UserController interface {
	ShowUsers() []entity.User
	ShowOwners() []entity.Owner
	QueryOwners1(c *gin.Context) entity.Owner
	Add(c *gin.Context) entity.User
	Update(c *gin.Context) entity.User
	Delete(c *gin.Context) entity.User
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

func (c controller) ShowOwners() []entity.Owner {
	return c.service.ShowOwners()
}

func (c controller) Add(ctx *gin.Context) entity.User {
	var user entity.User
	ctx.BindJSON(&user)
	c.service.Add(user)
	return user
}
func (c controller) QueryOwners1(ctx *gin.Context) entity.Owner {
	var owner entity.Owner
	var owner1 entity.Owner

	id, _ := strconv.ParseInt(ctx.Query("id"), 10, 32)
	log.Println("ID:", id)
	owner.Id = int(id)
	owner.Name = ""
	owner.Books = nil
	owner1 = c.service.QueryOwners1(owner)
	return owner1
}

func QueryString() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user entity.User
		user.ID = ctx.Query("id")
		user.Name = ctx.Query("name")
		ctx.JSON(200, user)
	}
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

func (c controller) ParamString(ctx *gin.Context) entity.User {
	var user entity.User
	user.ID = ctx.Param("id")
	user.Name = ctx.Param("name")
	c.service.Add(user)
	return user
}
