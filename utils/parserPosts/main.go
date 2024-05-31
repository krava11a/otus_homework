package main

import (
	"bufio"
	"bytes"
	"database/sql"
	"fmt"
	"strings"

	"os"

	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 15432
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

func main() {
	fmt.Println("### Read as reader ###")
	f, err := os.Open("posts.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Чтение файла с ридером
	wr := bytes.Buffer{}
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		wr.WriteString(sc.Text())
	}
	ps := strings.Split(wr.String(), ". ")

	InitDB()
	OpenConnectionDb()

	txn, err := Db.Begin()
	CheckError(err)

	stmt, err := txn.Prepare(pq.CopyIn("debug_posts", "text"))
	CheckError(err)

	for _, p := range ps {
		_, err = stmt.Exec(p)
		CheckError(err)
		// fmt.Println(p)
	}

	_, err = stmt.Exec()
	CheckError(err)

	err = stmt.Close()
	CheckError(err)

	err = txn.Commit()
	CheckError(err)

}
