package services

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "example"
	dbname   = "otus_homework"
)

var Db *sql.DB

func InitDB() {
	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// open database
	database, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	// close database
	defer database.Close()

	// check database
	err = database.Ping()
	CheckError(err)

	fmt.Println("Connected!")
}

func OpenConnectionDb() error {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		fmt.Println(err)
		return err
	}
	Db = db
	return nil
}

func CloseConnectionDb() {
	Db.Close()
}

func CheckError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
