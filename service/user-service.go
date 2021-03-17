package service

import (
	"log"

	"github.com/shubhamsnehi/gin-demo/database"
	"github.com/shubhamsnehi/gin-demo/entity"
)

var err error

type UserService interface {
	Add(entity.User) entity.User
	Update(entity.User)
	Delete(entity.User)
	ShowUsers() []entity.User
	ShowOwners() []entity.Owner
	QueryOwners1(entity.Owner) entity.Owner
}

//userService s
type userService struct {
}

func New() UserService {
	if err = database.Open(); err != nil {
		log.Println("Unsucessful")
	} else {
		log.Println("DB Connected Sucessfully")
	}
	return &userService{}
}

func (service *userService) Add(user entity.User) entity.User {
	database.DB.Create(&user)
	return user
}
func (service *userService) Update(user entity.User) {
	database.DB.Save(&user)
}

func (service *userService) Delete(user entity.User) {
	database.DB.Delete(&user, user.ID)
}

func (service *userService) ShowUsers() []entity.User {
	user := []entity.User{}
	database.DB.Table("usertbs").Find(&user)
	return user
}

func (service *userService) ShowOwners() []entity.Owner {
	result := []entity.Owner{}
	books := []entity.Book{}
	database.DB.Find(&result)
	for i := 0; i < len(result); i++ {
		database.DB.Table("books").Where("owner_id = ?", i+1).Scan(&books)
		result[i].Books = books
	}
	return result
}

func (service *userService) QueryOwners1(owner entity.Owner) entity.Owner {
	result := entity.Owner{}
	books := []entity.Book{}
	database.DB.Find(&result, owner.Id)
	database.DB.Table("books").Where("owner_id = ?", owner.Id).Scan(&books)
	result.Books = books
	return result
}
