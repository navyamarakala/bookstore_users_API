package user_db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

/*const (
	userName = os.Getenv(USERDB_USERNAME)
	password = os.Getenv(USERDB_PASSWORD)
	dbName   = os.Getenv(USERDB_DBNAME)
)*/

var (
	Db *sql.DB
)

func init() {
	dataSource := fmt.Sprintf("%s:%s@/%s?parseTime=true", "root", "root", "users_db")
	var err error
	Db, err = sql.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	if err := Db.Ping(); err != nil {
		panic(err)
	}
	log.Println("Connected to DB Sucessfully")
}
