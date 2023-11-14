package main

import (
	"fmt"
	"time"

	"github.com/bwmarrin/snowflake"
)

func GenerateOneNum(node *snowflake.Node) {
	fmt.Println("----------------------------------")

	// Generate a snowflake ID.
	id := node.Generate()

	// Print out the ID in a few different ways.
	fmt.Printf("Int64  ID: %d\n", id)
	fmt.Printf("String ID: %s\n", id)
	fmt.Printf("Base2  ID: %s\n", id.Base2())
	fmt.Printf("Base64 ID: %s\n", id.Base64())

	// Print out the ID's timestamp
	// id.Time() 为毫秒，需要使用 time.Unix(秒,纳秒) 转化，将 ms 转化为 ns(纳秒)
	fmt.Printf("ID Time  : %d, ID human Time : %+v\n", id.Time(), time.Unix(0, id.Time()*int64(time.Millisecond)))

	// Print out the ID's node number
	fmt.Printf("ID Node  : %d\n", id.Node())

	// Print out the ID's sequence number
	fmt.Printf("ID Step  : %d\n", id.Step())

	// Generate and print, all in one.
	fmt.Printf("ID       : %d\n", node.Generate().Int64())
}

func main() {
	// Create a new Node with a Node number of 1
	node, err := snowflake.NewNode(2)
	if err != nil {
		fmt.Println(err)
		return
	}

	for i := 0; i < 10; i++ {
		GenerateOneNum(node)
	}

}
