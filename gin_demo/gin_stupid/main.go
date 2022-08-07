package main

import (
	"future-go/gin_demo/gin_stupid/router"
)

func main() {
	r := router.SetupRouter()
	r.Run()
}
