package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func main()  {
	dsName := "root:casa123@tcp(127.0.0.1:3306)/po0?" +
		"charset=utf8&parseTime=true"
	db, err:=sql.Open("mysql", dsName)
	if err != nil {
		fmt.Println(err)
	}
	db.SetMaxIdleConns(2)
	db.SetMaxOpenConns(5)
	db.SetConnMaxLifetime(7*time.Hour)
	defer db.Close()

	fmt.Println(db.Query("select now()"))
}