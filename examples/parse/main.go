package main

import (
	"fmt"
	"github.com/qwerty2586/nginx-marshal/nginx"
	"os"
)

func main() {
	b, err := os.ReadFile("test.conf")
	if err != nil {
		fmt.Printf("Opening file end with error: %s\n", err)
		return
	}
	conf := nginx.Parse(string(b))
	fmt.Println("Parsed nginx configuration:")
	fmt.Println(conf)
}
