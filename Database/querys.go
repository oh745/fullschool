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
	stmt,err := db.Prepare("SELECT id,title,status FROM todos")
	
	if err != nil{

	}

	rows,err := stmt.Query()
	if err != nil{

	}

	tlist := []Todo{}

	for rows.Next(){

		t := Todo{}
		err1 := rows.Scan(&t.ID,&t.Title, &t.Status)
		if err1 != nil{
			log.Fatal("error scan", err.Error())
		}
		tlist = append(tlist,t)
	}
	

	fmt.Println(tlist)
}