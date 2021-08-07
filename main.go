package main

import (
	"context"
	"database/sql"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

const targetRowCount int = 100_000_000
const rowsPerBatch int = 50

func main() {
	os.Remove("./test.db")
	db, errConn := sql.Open("sqlite3", "./test.db")
	checkErr(errConn)

	dbSetPragma(db)
	dbCreateTable(db)

	//insertPrepared(db)
	insertTransaction(db)
	print("\n")
}

func dbSetPragma(db *sql.DB) {
	db.Exec("PRAGMA journal_mode = OFF;")
	db.Exec("PRAGMA synchronous = OFF;")
	db.Exec("PRAGMA temp_store = MEMORY;")
	db.Exec("PRAGMA cache_size = 1000000;")
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

	for i := 0; i < rowsPerBatch; i++ {
		tuples = append(tuples, "b", "b", "5")
		baseString += "(?, ?, ?),"
	}

	_, err := db.Exec("BEGIN")
	checkErr(err)

	stmt, err := db.Prepare(baseString[:len(baseString)-1])
	checkErr(err)
	defer stmt.Close()

	var counter = 0
	var iterations = targetRowCount / rowsPerBatch
	for i := 0; i < iterations; i++ {
		_, err = stmt.Exec(tuples...)
		checkErr(err)

		counter += rowsPerBatch
		if (counter % 100_000) == 0 {
			print("#")
		}
	}

	_, err = db.Exec("COMMIT")
	checkErr(err)
}

func insertTransaction(db *sql.DB) {
	var baseString = "INSERT into userinfo (username, departname, created) values"
	tuples := []interface{}{}

	for i := 0; i < rowsPerBatch; i++ {
		tuples = append(tuples, "b", "b", "5")
		baseString += "(?, ?, ?),"
	}

	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	checkErr(err)

	stmt, err := tx.PrepareContext(ctx, baseString[:len(baseString)-1])
	checkErr(err)
	defer stmt.Close()

	var counter = 0
	var iterations = targetRowCount / rowsPerBatch
	for i := 0; i < iterations; i++ {
		_, err = stmt.Exec(tuples...)
		checkErr(err)

		counter += rowsPerBatch
		if (counter % 100_000) == 0 {
			print("#")
		}
	}

	err = tx.Commit()
	checkErr(err)
}
