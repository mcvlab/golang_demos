package main

import (
	"embed"
	"fmt"
)

//go:embed assets/example.txt
var fs embed.FS

func main() {
	data, _ := fs.ReadFile("assets/example.txt")
	fmt.Println(string(data))
}
