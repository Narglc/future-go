package main

import (
	. "future-go/gin_demo/gin_todo/router"
)

func main() {
	r := InitRouter()
	r.Run()
}
