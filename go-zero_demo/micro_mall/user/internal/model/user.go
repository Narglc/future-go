package model

User struct {
	Id     int64  `db:"id"`
	Name   string `db:"name"`
	Gender string `db:"gender"`
}