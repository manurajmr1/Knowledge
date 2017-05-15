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
	ch := make(chan int, 50)
	start := time.Now()
	var count = 200
	// runtime.GOMAXPROCS(2)

	for i := 1; i <= count; i++ {
		go selectQuery(db, ch, i)
		// <-ch
	}

	// time.Sleep(time.Second * 2)

	for i := 1; i <= count; i++ {
		<-ch
	}
	// fmt.Println(runtime.NumCPU())
	close(ch)
	db.Close()
	elapsed := time.Since(start)
	fmt.Println(elapsed)
}

func selectQuery(db *fastsql.DB, ch chan int, cnt int) {

	// rows, err := db.Query("select u.* from users u left join address a on u.id = a.u_id")
	stmt, err := db.Prepare("select u.* from users u left join address a on u.id = a.u_id   WHERE u.id >  ?")
	defer stmt.Close()

	rows, err := stmt.Query(cnt)
	defer rows.Close()
	// rows, err := db.Query("SELECT SLEEP(2)")
	checkErr(err)
	// fmt.Println(cnt)

	// go func() {
	// 	for rows.Next() {
	// 		var id int
	// 		var user_fname string
	// 		var user_lname string
	// 		err = rows.Scan(&id, &user_fname, &user_lname)
	// 		checkErr(err)
	// 		fmt.Println(cnt, id, user_fname, user_lname)
	// 		// fmt.Println(user_fname)
	// 		// fmt.Println(user_lname)
	// 	}
	// 	ch <- "inside"
	// }()
	// var temp string
	for rows.Next() {
		var id int
		var user_fname string
		var user_lname string
		err = rows.Scan(&id, &user_fname, &user_lname)
		// checkErr(err)
		// fmt.Println(cnt, id, user_fname, user_lname)
		// temp = user_fname + user_lname
		// fmt.Println(user_fname)
		// fmt.Println(user_lname)
	}

	// fmt.Println(rows)
	ch <- 1

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
