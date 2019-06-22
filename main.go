package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	r.GET("/ping", pingHandler)
	r.POST("/ping", pingPost)
	r.POST("/getstudent", postStudentHandler)
	r.GET("/getstudent", getStudent)
	r.Run()
}

func pingHandler(c *gin.Context) {
	response := gin.H{
		"message": "pong",
	}
	c.JSON(200, response)
}

func pingPost(c *gin.Context) {
	response := gin.H{
		"message": "pong tai",
	}
	c.JSON(200, response)
}

func getStudent(c *gin.Context) {
	student := []Student{}
	for _, s := range student {
		student = append(student, s)
	}
	// ss := map[string]Student{
	// 	"12345": Student{Name: "oh"},
	// }

	c.JSON(http.StatusOK, student)
}

func postStudentHandler(c *gin.Context) {
	s := Student{}
	fmt.Printf("before bind %v\n", s)
	if err := c.ShouldBindJSON(&s); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	fmt.Printf("after bind %v\n", s)
}

type Student struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}
