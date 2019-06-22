package main

import(
	"database/sql"
	"log"
	"os"
	_"github.com/lib/pq"
)


func main(){
	//uri := "postgres://pjsgdumf:HuQ36jnSN6_X-iM9sXvEFfsp9G96Wkyh@satao.db.elephantsql.com:5432/pjsgdumf"
	uri := os.Getenv("DATABASE_URL")
	db,err := sql.Open("postgres",uri)
	if err != nil {
		log.Fatal("fatal",err.Error())
	}
	defer db.Close()

	title := "Home Work"
	status := "Inactive"
	query := `INSERT INTO todos (title,status) VALUES ($1,$2) RETURNING id`
	var id int
	row := db.QueryRow(query ,title,status)
	err = row.Scan(&id)
	if err != nil {
		log.Fatal("cant scan id",err.Error())
	}
}