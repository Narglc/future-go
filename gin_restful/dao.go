package main

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func CreateUser(c *gin.Context) {

}

func FetchAllUsers(c *gin.Context) {

}

func UpdateUser(c *gin.Context) {

}

func DeleteUser(c *gin.Context) {

}

func FetchSingleUser(c *gin.Context) {
	id := c.Param("id")

	db, err := sql.Open("mysql", "root:MySql1990328@(127.0.0.1:3306)/gin_restful_api?charset=utf8")
	checkErr(err)

	defer db.Close()

	err = db.Ping()
	checkErr(err)

	var (
		person Person
		result gin.H
	)
	row := db.QueryRow("select id, name from user_info where id = ?;", id)
	err = row.Scan(&person.Id, &person.Name)
	if err != nil {
		// If no results send null
		result = gin.H{
			"result": nil,
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": person,
			"count":  1,
		}
	}
	c.JSON(http.StatusOK, result)
}
