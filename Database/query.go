package main

import(
	"database/sql"
	"log"
	"fmt"
	"os"
	_"github.com/lib/pq"
)

type Todo struct {
	ID int
	Title string
	Status string
}


func main(){
	db,_ := sql.Open("postgres",os.Getenv("DATABASE_URL"))
	stmt,err := db.Prepare("SELECT id,title,status FROM todos WHERE id = $1")
	
	if err != nil{

	}

	row := stmt.QueryRow(1)

	t := Todo{}

	err1 := row.Scan(&t.ID,&t.Title, &t.Status)
	if err1 != nil{
		log.Fatal("error scan", err.Error())
	}

	fmt.Println(t)
}