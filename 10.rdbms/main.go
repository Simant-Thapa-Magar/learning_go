package main

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func index(w http.ResponseWriter, req *http.Request) {
	_, err := io.WriteString(w, "Connected successfully")
	checkError(err)
}

func createTable(w http.ResponseWriter, req *http.Request) {
	var st, count string
	var c int
	tableName := req.FormValue("table_name")

	if tableName == "" {
		fmt.Fprintln(w, "Need some table name")
		return
	}

	st = `SELECT COUNT(*) FROM information_schema.tables WHERE table_schema = DATABASE() AND table_name = "` + tableName + `"`

	r, err := db.Query(st)

	checkError(err)

	for r.Next() {
		err = r.Scan(&count)
		checkError(err)
		c, err = strconv.Atoi(count)
		checkError(err)

		if c > 0 {
			fmt.Fprintln(w, "Table with name "+tableName+" already exists !!!")
			return
		}
	}

	st = `CREATE TABLE IF NOT EXISTS ` + tableName + ` (id INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY, data VARCHAR(50))`

	stmt, err := db.Prepare(st)

	checkError(err)

	result, err := stmt.Exec()

	checkError(err)

	n, err := result.RowsAffected()
	checkError(err)

	fmt.Fprintln(w, "Created table "+tableName+" successfully", n)

}

func insertData(w http.ResponseWriter, req *http.Request) {
	var c string
	var count int
	tableName := req.FormValue("table_name")
	dataName := req.FormValue("data")

	if tableName == "" || dataName == "" {
		fmt.Fprintln(w, "Missing table name or data")
		return
	}

	st := `SELECT COUNT(*) FROM information_schema.tables WHERE table_schema = DATABASE() AND table_name = "` + tableName + `"`

	row, err := db.Query(st)

	checkError(err)

	for row.Next() {
		err = row.Scan(&c)
		checkError(err)

		count, err = strconv.Atoi(c)
		checkError(err)

		if count == 0 {
			fmt.Fprintln(w, "No table with name "+tableName+" exists")
			return
		}
	}

	st = `INSERT INTO ` + tableName + ` (data) VALUES("` + dataName + `");`

	pStmt, err := db.Prepare(st)

	checkError(err)

	result, err := pStmt.Exec()

	checkError(err)

	rowEffected, err := result.RowsAffected()

	checkError(err)

	if rowEffected > 0 {
		id, err := result.LastInsertId()
		checkError(err)
		fmt.Fprintln(w, "Data created successfully with id ", id)
		return
	}

	fmt.Fprintln(w, "Error inserting data")
}

func readData(w http.ResponseWriter, req *http.Request) {
	var c, d string
	var count int
	tableName := req.FormValue("table_name")

	if tableName == "" {
		fmt.Fprintln(w, "Table name is required")
		return
	}

	st := `SELECT COUNT(*) FROM information_schema.tables WHERE table_schema = DATABASE() AND table_name = "` + tableName + `"`

	rows, err := db.Query(st)

	checkError(err)

	for rows.Next() {
		err = rows.Scan(&c)
		checkError(err)
		count, err = strconv.Atoi(c)
		checkError(err)

		if count == 0 {
			fmt.Fprintln(w, "No table with name "+tableName+" exists")
			return
		}
	}

	rows, err = db.Query(`SELECT data FROM ` + tableName)

	checkError(err)

	data := "Data retrieved:\n"

	for rows.Next() {
		err = rows.Scan(&d)
		checkError(err)
		data += d + "\n"
	}

	fmt.Fprintln(w, data)
}

func checkError(e error) {
	if e != nil {
		log.Panic(err)
	}
}

func main() {
	db, err = sql.Open("mysql", "go:lang@tcp(127.0.0.1:3306)/golang")

	checkError(err)

	defer db.Close()

	err = db.Ping()

	checkError(err)

	http.HandleFunc("/", index)
	http.HandleFunc("/create_table", createTable)
	http.HandleFunc("/insert_data", insertData)
	http.HandleFunc("/read_data", readData)
	http.ListenAndServe(":8080", nil)
}
