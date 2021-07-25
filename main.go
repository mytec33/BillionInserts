package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

const valuesCount int = 10

func main() {
	db, errConn := sql.Open("sqlite3", "./test.db")
	checkErr(errConn)

	dbSetPragma(db)
	dbCreateTable(db)

	stmt, err := db.Prepare("INSERT INTO userinfo(username, departname, created) values(?, ?, ?),(?, ?, ?),(?, ?, ?),(?, ?, ?),(?,?,?),(?, ?, ?),(?, ?, ?),(?, ?, ?),(?, ?, ?),(?, ?, ?)")
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
	_, err := stmt.Exec("bill", "研发部门", "2021-07-25", "bill", "研发部门", "2021-07-25",
		"bill", "研发部门", "2021-07-25", "bill", "研发部门", "2021-07-25",
		"bill", "研发部门", "2021-07-25", "bill", "研发部门", "2021-07-25",
		"bill", "研发部门", "2021-07-25", "bill", "研发部门", "2021-07-25",
		"bill", "研发部门", "2021-07-25", "bill", "研发部门", "2021-07-25")
	checkErr(err)
}
