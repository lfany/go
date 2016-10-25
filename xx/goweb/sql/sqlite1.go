package main

import (
	_ "github.com/mattn/go-sqlite3"
	"database/sql"
	"fmt"
)

func main() {
	db, err := sql.Open("sqlite3", "./test.db")
	checkError(err)
	defer db.Close()

	// 插入数据
	stmt, err := db.Prepare("INSERT INTO userinfo(username, departname,created) values(?,?,?)")

	res, err := stmt.Exec("abc", "Android开发", "2012-12-09")
	checkError(err)

	id, err := res.LastInsertId()
	checkError(err)

	fmt.Println(id)
//	更新数据
	stmt, err = db.Prepare("update userinfo set username=? where uid=?")
	checkError(err)

	res, err = stmt.Exec("ABC", id)
	checkError(err)

	affect, err := res.RowsAffected()
	checkError(err)

	fmt.Println(affect)

//	查询数据
	rows, err := db.Query("select * from userinfo")
	checkError(err)

	for rows.Next() {
		var uid int
		var username string
		var department string
		var created string
		err = rows.Scan(&uid, &username, &department, &created)
		checkError(err)
		fmt.Println(uid, "\t", username, "\t", department, "\t", created)
	}

//	删除数据
	stmt, err = db.Prepare("delete from userinfo where uid = ?")
	checkError(err)

	res, err = stmt.Exec(id)
	checkError(err)

	affect, err = res.RowsAffected()
	checkError(err)

	fmt.Println(affect)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}