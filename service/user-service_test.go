package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/shubhamsnehi/REST-Api-with-Golang/entity"
	"github.com/stretchr/testify/assert"
)

func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

type User1 struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func TestShowUsers(t *testing.T) {
	gin.SetMode(gin.TestMode)
	handler := New().ShowUsers()
	router := gin.Default()
	router.GET("/users", func(ctx *gin.Context) {
		ctx.JSON(200, handler)
	})

	req, err := http.NewRequest("GET", "/users", nil)
	if err != nil {
		fmt.Println(err)
	}
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(t, resp.Code, 200)
	expected := `[{"id":"4","name":"S R Snehi"},{"id":"23","name":"Test"},{"id":"35","name":"Shubham"},{"id":"45","name":"Heyy"},{"id":"46","name":"Heyy"}]`
	assert.Equal(t, resp.Body.String(), expected)
}

func TestAdd(t *testing.T) {
	gin.SetMode(gin.TestMode)
	handler := New().Add(entity.User{Name: "Heyy"})
	router := gin.Default()
	router.POST("/users", func(ctx *gin.Context) {
		ctx.JSON(200, handler)
	})

	req, err := http.NewRequest("POST", "/users", nil)
	if err != nil {
		fmt.Println(err)
	}
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(t, resp.Code, 200)
	expected := "Heyy"
	// log.Println(expected)
	var response User1
	err = json.Unmarshal([]byte(resp.Body.String()), &response)
	// log.Println(resp.Body)
	// log.Println(response.Name)
	assert.Equal(t, response.Name, expected)
}

// func TestUpdate(t *testing.T) {
// 	// user := User1{ID: "1", Name: "Heyyy"}
// 	gin.SetMode(gin.TestMode)
// 	handler := New().Update(entity.User{Name: "Heyyy"})
// 	router := gin.Default()
// 	router.PATCH("/users", func(ctx *gin.Context) {
// 		ctx.JSON(200, handler)
// 	})

// 	req, err := http.NewRequest("PATCH", "/users", nil)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	resp := httptest.NewRecorder()
// 	router.ServeHTTP(resp, req)
// 	assert.Equal(t, resp.Code, 200)
// 	expected := "Heyyy"
// 	// log.Println(expected)
// 	var response User1
// 	err = json.Unmarshal([]byte(resp.Body.String()), &response)
// 	// log.Println(resp.Body)
// 	// log.Println(response.Name)
// 	assert.Equal(t, response.Name, expected)
// }
