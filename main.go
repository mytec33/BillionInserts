package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

const valuesCount int = 100

func main() {
	db, errConn := sql.Open("sqlite3", "./test.db")
	checkErr(errConn)

	dbSetPragma(db)
	dbCreateTable(db)

	insertPrepared(db)
	print("\n")
}

func dbSetPragma(db *sql.DB) {
	db.Exec("PRAGMA journal_mode = OFF;")
	db.Exec("PRAGMA synchronous = OFF;")
	db.Exec("PRAGMA temp_store = MEMORY;")
	//db.Exec("PRAGMA cache_size = 32768;")
	//db.Exec("PRAGMA page_size = 16384")
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

func insertPrepared(db *sql.DB) {
	var baseString = "INSERT into userinfo (username, departname, created) values"
	tuples := []interface{}{}

	for i := 0; i < valuesCount; i++ {
		tuples = append(tuples, "bill", "研发部门", "2021-07-25")
		baseString += "(?, ?, ?),"
	}

	stmt, err := db.Prepare(baseString[:len(baseString)-1])
	checkErr(err)
	defer stmt.Close()

	_, err = db.Exec("BEGIN")
	checkErr(err)

	var counter = 0
	var rows = 100000000 / valuesCount
	for i := 0; i < rows; i++ {
		_, err = stmt.Exec(tuples...)
		checkErr(err)

		counter += valuesCount
		if (counter % 100000) == 0 {
			print("#")
		}
	}

	_, err = db.Exec("COMMIT")
	checkErr(err)
}
