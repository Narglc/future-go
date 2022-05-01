package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

var DB *sqlx.DB

type User struct {
	Id       int    `db:"id"`
	Name     string `db:"name"`
	Email    string `db:"email"`
	Password string `db:"password"`
}

// SQL 建表
//CREATE TABLE `users` (
//`id` INT(10) PRIMARY KEY AUTO_INCREMENT,
//`name` VARCHAR(50) NULL DEFAULT NULL COLLATE 'utf8_general_ci',
//`email` VARCHAR(128) NULL DEFAULT NULL COLLATE 'utf8_general_ci',
//`password` VARCHAR(128) NULL DEFAULT NULL COLLATE 'utf8_general_ci'
//)

func init() {
	db, err := sqlx.Open("mysql", "root:MySql1990328@(127.0.0.1:3306)/blog")
	if err != nil {
		fmt.Printf("open mysql err:", err)
		return
	}
	DB = db
}

// 增加
func insertDemo() {
	res, err := DB.Exec("INSERT INTO users (name, email, password) VALUES(?,?,?)", "mysqldemo", "my@12.com", "11111")
	if err != nil {
		fmt.Println("insert err:", err)
		return
	}
	id, err := res.LastInsertId()
	if err != nil {
		fmt.Println("get last id err", err)
		return
	}
	fmt.Println("insert success:", id)
}

// 修改
func updateDemo() {
	_, err := DB.Exec("UPDATE users SET name=? WHERE id=?", "mysqlupdate", 1)
	if err != nil {
		fmt.Println("update err:", err)
		return
	}
	fmt.Println("update succ.")
}

// 删除
func delDemo() {
	_, err := DB.Exec("DELETE FROM users WHERE id = ?", 2)
	if err != nil {
		fmt.Println("del err:", err)
		return
	}
	fmt.Println("del succ")
}

// 查询
func queryDemo() {
	var users []User
	err := DB.Select(&users, "SELECT * FROM users WHERE id = ?", 3)
	if err != nil {
		fmt.Println("query err:", err)
		return
	}
	fmt.Println("select succ and result:", users)
}

func queryAllDataDemo() {
	person := []User{}
	DB.Select(&person, "SELECT name,email,password FROM users")
	// 逐行打印
	fmt.Println("--------------- all data is:------------")
	for idx, each := range person {
		fmt.Printf("idx:%d, name:%v, email:%v, password:%v\n", idx, each.Name, each.Email, each.Password)
	}
}

func transDemo() {
	log.Println("begin transid...")
	tx, err := DB.Begin()
	if err != nil {
		log.Println("DB begin failed err:", err)
		return
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		}
	}()

	// 此处写一个demo
	if _, err = tx.Exec("INSERT INTO users (name,email,password) VALUES(?,?,?);", "demo2", "demo2@123.com", "123556"); err != nil {
		log.Println("trans err:", err)
		tx.Rollback()
		return
	}

	if _, err = tx.Exec("INSERT INTO users (name,email,password) VALUES(?,?,?);", "demo3", "demo3@123.com", "444"); err != nil {
		log.Println("trans err:", err)
		tx.Rollback()
		return
	}

	err = tx.Commit()
	return
}

func main() {
	insertDemo()
	updateDemo()
	delDemo()
	queryDemo()
	queryAllDataDemo()
	transDemo()
}
