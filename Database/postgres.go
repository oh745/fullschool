package main

import(
	"database/sql"
	"fmt"
	"log"
	_"github.com/lib/pq"
)

func main(){
	uri := "postgres://pjsgdumf:HuQ36jnSN6_X-iM9sXvEFfsp9G96Wkyh@satao.db.elephantsql.com:5432/pjsgdumf"
	db,err := sql.Open("postgres",uri)
	if err != nil {
		log.Fatal("fatal",err.Error())
	}
	defer db.Close()

	createTb := 
	`CREATE TABLE IF NOT EXISTS NAJA(
		id SERIAL PRIMARY KEY,
		title TEXT,
		status TEXT
	);`
	_,err = db.Exec(createTb)
	if err != nil {
		log.Fatal("cant create table",err.Error())
	}

	fmt.Println("oo")

}