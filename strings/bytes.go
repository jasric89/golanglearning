package main

/*
This program uses the Hashprefix function thats available in both Strings and Bytes Package.
HasPrefix in strings checks the prefix regardless on how many bytes are used by the characters.
HasPrefix in bytes determines if the price variable begins with a specific sequence of bytes. Its useful to use this package to store binary data in memory.
*/
import (
	"bytes"
	"fmt"
	"strings"
)

func main() {

	price := "£100"
	fmt.Println("Strings Prefix:", strings.HasPrefix(price, "£"))
	fmt.Println("Bytes Prefix:", bytes.HasPrefix([]byte(price), []byte{226, 130}))
}
