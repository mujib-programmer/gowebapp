package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"time"
)

func main() {
	db, err := sql.Open("sqlite3", "./go-sqlite3.db")
	checkErr(err)

	// insert
	stmt, err := db.Prepare("INSERT INTO userinfo (username, departname, created) values(?,?,?)")
	checkErr(err)

	res, err := stmt.Exec("mujib", "software coder", "2018-08-29")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println(id)

	// update
	stmt, err = db.Prepare("update userinfo set username=? where uid=?")
	checkErr(err)

	res, err = stmt.Exec("mujibur rochman", id)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	// query
	rows, err := db.Query("SELECT * FROM userinfo")
	checkErr(err)

	for rows.Next() {
		var uid int
		var username string
		var departname string
		var created time.Time

		err = rows.Scan(&uid, &username, &departname, &created)
		checkErr(err)

		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(departname)
		fmt.Println(created)

	}

	rows.Close() // good habit to close

	// delete
	stmt, err = db.Prepare("DELETE FROM userinfo WHERE uid=?")
	checkErr(err)

	res, err = stmt.Exec(id)
	checkErr(err)

	affect, err = res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)
	db.Close()

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
