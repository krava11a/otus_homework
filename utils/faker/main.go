package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strings"

	"github.com/lib/pq"
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

func getSqlInsertFromCsv(csvFileName string) {
	file, err := os.Open(csvFileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	txn, err := Db.Begin()
	CheckError(err)

	stmt, err := txn.Prepare(pq.CopyIn("users", "first_name", "second_name", "biography", "birthdate", "city", "hp"))
	CheckError(err)

	record := make([]string, 3)
	fio := make([]string, 2)

	for scanner.Scan() {
		record = strings.Split(scanner.Text(), ",")
		fio = strings.Split(record[0], " ")
		_, err = stmt.Exec(fio[1], fio[0], "homework2,jmeter", record[1], record[2], "askdjflkajsdlfkja;lfd")
		CheckError(err)
	}

	_, err = stmt.Exec()
	CheckError(err)

	err = stmt.Close()
	CheckError(err)

	err = txn.Commit()
	CheckError(err)

}

func main() {
	InitDB()
	OpenConnectionDb()
	getSqlInsertFromCsv("people.v2.csv")
}
