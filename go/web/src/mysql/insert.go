package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:my$ql@/golang?charset=utf8")
	checkErr(err)
	ch := make(chan int, 5)
	start := time.Now()
	var count = 100

	go insertquery(db, count, ch)

	<-ch

	db.Close()
	elapsed := time.Since(start)
	fmt.Println(elapsed)
}

func insertquery(db *sql.DB, cnt int, ch chan int) {

	stmtIns, err := db.Prepare("INSERT  user  VALUES( ?, ?, ?, ? )") // ? = placeholder
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer stmtIns.Close()
	for i := 1; i < cnt; i++ {
		_, err = stmtIns.Exec(cnt, "fname", "lname", "addresss") // Insert tuples (i, i^2)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		fmt.Println(i)
	}

	ch <- 1
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
