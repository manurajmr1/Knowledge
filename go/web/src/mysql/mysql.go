package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func main() {
	db, err := sql.Open("mysql", "root:my$ql@/golang?charset=utf8")
	checkErr(err)
	ch := make(chan string)
	start := time.Now()
	var count = 4
	for i := 1; i <= count; i++ {
		go selectQuery(db, ch, i)
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

func selectQuery(db *sql.DB, ch chan string, cnt int) {
	// rows, err := db.Query("select u.* from users u left join address a on u.id = a.u_id")
	stmt, err := db.Prepare("select u.* from users u left join address a on u.id = a.u_id WHERE u.id > ?")
	defer stmt.Close()

	rows, err := stmt.Query(cnt)
	defer rows.Close()
	// rows, err := db.Query("SELECT SLEEP(2)")
	checkErr(err)
	// fmt.Println(cnt)
	for rows.Next() {
		var id int
		var user_fname string
		var user_lname string
		err = rows.Scan(&id, &user_fname, &user_lname)
		checkErr(err)
		fmt.Println(cnt, id, user_fname, user_lname)
		// fmt.Println(user_fname)
		// fmt.Println(user_lname)

	}
	// fmt.Println(cnt)
	ch <- "got"

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
