package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/rmulley/go-fast-sql"
	"time"
)

func main() {
	db, err := fastsql.Open("mysql", "root:my$ql@/golang?charset=utf8", 100)
	checkErr(err)
	ch := make(chan string)
	start := time.Now()
	var count = 1
	for i := 1; i <= count; i++ {
		go selectQuery(db, ch, 100)
		// <-ch
	}

	for i := 1; i <= count; i++ {
		<-ch
	}
	close(ch)
	db.Close()
	elapsed := time.Since(start)
	fmt.Println(elapsed)
}

func selectQuery(db *fastsql.DB, ch chan string, cnt int) {
	var (
		err error
	)
	// stmtIns, err := db.Prepare("INSERT  user  VALUES( ?, ?, ?, ? )") // ? = placeholder
	// if err != nil {
	// 	panic(err.Error()) // proper error handling instead of panic in your app
	// }
	// defer stmtIns.Close()
	for i := 1; i < cnt; i++ {
		// _, err = stmtIns.Exec(cnt, "fname", "lname", "addresss") // Insert tuples (i, i^2)
		// if err != nil {
		// 	panic(err.Error()) // proper error handling instead of panic in your app
		// }

		if err = db.BatchInsert("INSERT INTO user VALUES(?, ?, ?, ?);", i, "fname", "lname", "addresss"); err != nil {
			fmt.Println(err)
		}
		fmt.Println(i)
	}

	// fmt.Println(cnt)
	ch <- "got"

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
