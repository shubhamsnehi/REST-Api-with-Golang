package service

import (
	"log"

	"github.com/shubhamsnehi/gin-demo/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

type UserService interface {
	Add(entity.User) entity.User
	Update(entity.User) 
	Delete(entity.User) 
	ShowUsers() []entity.User
}

//userService s
type userService struct {
	users []entity.User
}

func New() UserService {
	return &userService{}
}

func (service *userService) Add(user entity.User) entity.User {
	dsn := "root:@tcp(localhost)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{}) //db connection
	if err != nil {
		log.Println("Could not connect")
	} else {
		db.Create(&user)
	}
	return user
}
func (service *userService) Update(user entity.User) {
	dsn := "root:@tcp(localhost)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{}) //db connection
	if err != nil {
		log.Println("Could not connect")
	} else {
		db.Save(&user)
	}
}

func (service *userService) Delete(user entity.User) {
	dsn := "root:@tcp(localhost)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{}) //db connection
	if err != nil {
		log.Println("Could not connect")
	} else {
		db.Delete(&user, user.ID)
	}
}

func (service *userService) ShowUsers() []entity.User {
	user := []entity.User{}
	dsn := "root:@tcp(localhost)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{}) //db connection
	if err != nil {
		log.Println("Could not connect")
	} else {
		db.AutoMigrate(&entity.User{})
		// db.Create([]entity.User{{Name: "Golang"},{Name: "Python"}})
		db.Find(&user)
	}
	return user
}
