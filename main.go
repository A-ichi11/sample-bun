package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
)

type User struct {
	Id       int64  `bun:"id"`
	Name     string `bun:"name"`
	Age      int    `bun:"age"`
	Password string `bun:"password"`
}

func main() {

	engine, err := sql.Open("mysql", "root:root@tcp([127.0.0.1]:3306)/sample_db?charset=utf8mb4&parseTime=true")
	if err != nil {
		panic(err)
	}
	db := bun.NewDB(engine, mysqldialect.New())

	// 動かす時にコメントインする
	// dropTable(db)
	// createTable(db)
	// insertOne(db)
	// insertAll(db)
	// getOne(db)
	// getAll(db)
	// delete(db)
	update(db)
}

func dropTable(db *bun.DB) {
	_, err := db.NewDropTable().Model((*User)(nil)).Exec(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("dropTable")
}

func createTable(db *bun.DB) {
	_, err := db.NewCreateTable().Model((*User)(nil)).Exec(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println("createTable")
}

func insertOne(db *bun.DB) {
	user := User{
		Name:     "太郎",
		Password: "パスワード",
		Age:      20,
	}

	_, err := db.NewInsert().Model(&user).Exec(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println("insertOne")
}

func insertAll(db *bun.DB) {
	user2 := User{
		Name:     "花子",
		Password: "パスワード",
		Age:      25,
	}
	user3 := User{
		Name:     "りょう",
		Password: "パスワード",
		Age:      30,
	}
	users := []User{user2, user3}
	_, err := db.NewInsert().Model(&users).Exec(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println("insertAll")
}

func getOne(db *bun.DB) {
	user := User{}
	err := db.NewSelect().Model(&user).Where("id = 1").Scan(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println("getOne", user)
}

func getAll(db *bun.DB) {
	users := []User{}
	err := db.NewSelect().Model(&users).OrderExpr("id ASC").Scan(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println("getAll", users)
}

func delete(db *bun.DB) {
	user := User{}
	_, err := db.NewDelete().Model(&user).Where("id = 1").Exec(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println("delete")
}

func update(db *bun.DB) {
	user := User{}
	_, err := db.NewUpdate().Model(&user).Set("age = 40").Where("id = ?", 2).Exec(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println("update")
}
