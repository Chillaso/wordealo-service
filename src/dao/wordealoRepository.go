package dao

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" // Blank import calls the init function of the package
)

const (
	DB_HOST  = "localhost"
	DB_PORT  = 3306
	DATABASE = "wordealo"
)

var (
	db *sql.DB
)

// Function that connect to a mysql database at localhost with name wordealo, username root and password toor and return the connection using the standard library
func init() {
	conn, err := sql.Open("mysql", fmt.Sprintf("root:toor@tcp(%s:%d)/%s", DB_HOST, DB_PORT, DATABASE))
	if err != nil {
		panic(err.Error())
	}
	db = conn
}

func GetPossibleWords(query string) []string {
	rows, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}

	// Provide a slice of string with 15 as a capacity, maximum of results
	var result = make([]string, 0, 15)
	for rows.Next() {
		var word string
		err = rows.Scan(&word)
		if err != nil {
			panic(err.Error())
		}
		result = append(result, word)
	}
	return result
}
