package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

const valuesCount int = 200

func main() {
	db, errConn := sql.Open("sqlite3", "./test.db")
	checkErr(errConn)

	dbSetPragma(db)
	dbCreateTable(db)

	stmt, err := db.Prepare(createInsertString())
	checkErr(err)
	defer stmt.Close()

	var counter = 0
	var rows = 1000000 / valuesCount

	for i := 0; i < rows; i++ {
		insertPrepared(stmt)

		counter += valuesCount
		if (counter % 1000) == 0 {
			print("#")
		}
	}
	print("\n")
}

func createInsertString() string {
	var baseString = "INSERT into userinfo (username, departname, created) values"

	for i := 1; i <= valuesCount; i++ {
		baseString += "(?, ?, ?),"
	}

	return baseString[:len(baseString)-1]
}

func dbSetPragma(db *sql.DB) {
	db.Exec("PRAGMA journal_mode = OFF;")
	db.Exec("PRAGMA synchronous = OFF;")
	db.Exec("PRAGMA temp_store = MEMORY;")
	//db.Exec("PRAGMA cache_size = 1000000;")
}

func dbCreateTable(db *sql.DB) {
	db.Exec(`CREATE TABLE IF NOT EXISTS userinfo 
		(
			id INTEGER not null primary key
			,username CHAR(10)
			,departname char(10) not null
			,created date not null
		)`)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func insertPrepared(stmt *sql.Stmt) {
	tuples := []interface{}{}

	for i := 0; i < valuesCount; i++ {
		tuples = append(tuples, "bill", "研发部门", "2021-07-25")
	}
	_, err := stmt.Exec(tuples...)
	checkErr(err)
}
