package main

import (
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"database/sql"
)

func main() {
	db, err := sql.Open("mysql", "root:00@/test?charset=utf8")
	checkErr(err)

	stmt, err := db.Prepare("Insert userinfo SET username=?,departname=?, created=?")
	checkErr(err)

	res, err := stmt.Exec("astaxie", "研发部门", "2012-12-09")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println(id)

	
}
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
