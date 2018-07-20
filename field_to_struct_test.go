package frame

import (
	"database/sql"
	"fmt"
	"strings"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

type Test struct {
	UserId    int    `sql:"id"`
	UserPhone string `sql:"phone"`
	Email     string `sql:"email"`
	CreatedAt string `sql:"created_at"`
}

func init() {
	db, err := sql.Open("mysql", "root:L8n5Du--@unix(/tmp/mysql.sock)/cars?charset=utf8")
	if err != nil {
		fmt.Printf("Open db fail! %#v", err)
	}
	db.SetMaxOpenConns(1000)
	db.SetMaxIdleConns(1000)
}

func TestModel(t *testing.T) {
	t.Log(">>>>>>>TestModel")

	if db == nil {
		t.Log("db is nil.")
		return
	}

	testRows := FindAll("select * from car_users")
	testRow := Find("select * from car_users Limit 1")

	t.Log(testRows)
	t.Log(testRow)

	t.Log("====>End")
}

func FindAll(sqlStr string) []Test {

	t := Test{}
	ts := []Test{}
	fieldString, values := FieldToStruct(&t)

	rows, err := db.Query(strings.Replace(sqlStr, "*", fieldString, 1))
	defer rows.Close()
	if err != nil {
		fmt.Println("sql open error ", err)
		return nil
	}

	for rows.Next() {
		rows.Scan(values...)

		ts = append(ts, t)
	}

	return ts
}

func Find(sqlStr string) Test {

	t := Test{}
	fieldString, values := FieldToStruct(&t)
	db.QueryRow(strings.Replace(sqlStr, "*", fieldString, 1)).Scan(values...)

	return t
}
