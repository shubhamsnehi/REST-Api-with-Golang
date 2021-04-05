package service

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

type User []struct {
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
	expected := `[{"id":"4","name":"S R Snehi"},{"id":"15","name":"Updated"},{"id":"17","name":"Hello World"},{"id":"20","name":"Shubham Snehi"},{"id":"23","name":"Test"},{"id":"25","name":"Test1"},{"id":"35","name":"Shubham"}]`
	assert.Equal(t, resp.Body.String(), expected)
}
