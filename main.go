package main

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/uptrace/bun"
)

type User struct {
	Id       int64 `bun:"id"`
	Name     string
	Age      int
	Password string    `bun:"password"`
	Created  time.Time `bun:"created"`
	Updated  time.Time `bun:"updated"`
}

func main() {

	sqldb, err := sql.Open("mysql", "root:root@tcp([127.0.0.1]:3306)/sample_db?charset=utf8mb4&parseTime=true")
	if err != nil {
		panic(err)
	}
	db := bun.NewDB(sqldb, mysqldialect.New())

	insert(db)

}

func insert(db) {

}
