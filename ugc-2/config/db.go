package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB(dataSourceName string) *sql.DB {
	var err error
	DB, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		fmt.Println("Error while connecting to the database :", err)
		return nil
	}

	err = DB.Ping()
	if err != nil {
		fmt.Println("Error ping action :", err)
		return nil
	}
	fmt.Println("Successfully connected to the database")
	return DB
}
