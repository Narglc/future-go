package main

import (
	"encoding/hex"
	"fmt"

	uuid "github.com/satori/go.uuid"
)

func NewUUIDV1() string {
	return uuid.NewV1().String()
}

func NewUUIDV2(domain byte) string {
	return uuid.NewV2(domain).String()
}

func NewUUIDV4() string {
	return uuid.NewV4().String()
}

// 不使用UUID自带的String()方法，即去除-
func NewUUIDV4Simple() string {
	id := uuid.NewV4()
	buf := make([]byte, 32)
	hex.Encode(buf[:], id[:])
	return string(buf)
}

func main() {
	fmt.Println("UUIDV1:", NewUUIDV1())
	fmt.Println("UUIDV2:", NewUUIDV2('x'))
	fmt.Println("UUIDV4:", NewUUIDV4())
	fmt.Println("UUIDV4Simple{去除横杠-}:", NewUUIDV4Simple())

	fmt.Println("============官方用例=================")
	officeExample()
}

func officeExample() {
	// Creating UUID Version 4
	// panic on error
	u1 := uuid.Must(uuid.NewV4(), nil)
	fmt.Printf("UUIDv4: %s\n", u1)

	// or error handling
	u2 := uuid.NewV4()
	fmt.Printf("UUIDv4: %s\n", u2)

	// Parsing UUID from string input
	u2, err := uuid.FromString("6ba7b810-9dad-11d1-80b4-00c04fd430c8")
	if err != nil {
		fmt.Printf("Something went wrong: %s", err)
		return
	}
	fmt.Printf("Successfully parsed: %s", u2)
}
