package main

import (
	"fmt"
)

func main() {
	fmt.Println(normalizeURL("https://blog.boot.dev/path/"))
	fmt.Println(normalizeURL("https://blog.boot.dev/path"))
	fmt.Println(normalizeURL("http://blog.boot.dev/path/"))
	fmt.Println(normalizeURL("http://blog.boot.dev/path"))


}
