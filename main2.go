package main

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"database/sql"
	"log"
	"os"
	_"github.com/lib/pq"
)

func main() {
	r := gin.Default()
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	r.GET("/todos", getTodosHandler)
	r.POST("/ping", pingPost)
	r.POST("/getstudent", postStudentHandler)
	r.GET("/getstudent", getStudent)
	r.Run()
}

func getTodosHandler(c *gin.Context) {
	db,_ := sql.Open("postgres",os.Getenv("DATABASE_URL"))
	stmt,_ := db.Prepare("SELECT id,title,status FROM todos")
	rows,_ := stmt.Query()

	todos := []Todo{}

	for rows.Next(){

		t := Todo{}
		err1 := rows.Scan(&t.ID,&t.Title, &t.Status)
		if err1 != nil{
			log.Fatal("error scan", err1.Error())
		}
		todos = append(todos,t)
	}

	
	c.JSON(200, todos)
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

type Todo struct {
	ID int
	Title string
	Status string
}

type Student struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}
