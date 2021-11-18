package database

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"
)

func NewDB() *sql.DB {
	// TODO use vipper to parse config file and set env valiables.
	DBMS := "mysql"
	// dbHost := "localhost"
	dbHost := "127.0.0.1"
	dbPort := "3306"
	dbUser := "gosqldemouser"
	dbPass := "gosqldemopass"
	dbName := "gosqldemodb"
	val := url.Values{}
	val.Add("charset", "utf8mb4")
	val.Add("parseTime", "true")
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())
	dbConn, err := sql.Open(DBMS, dsn)

	if err != nil {
		log.Fatal(err)
	}
	// 二つ目は:=ではなく=
	err = dbConn.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return dbConn
}
