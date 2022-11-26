package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello world!")
	fmt.Println("Hello" + "world!")
	fmt.Println("Hello world!"[0])
	fmt.Println(string("Hello world!"[0]))
	var s string = "Hello world!"

	fmt.Println(strings.Replace(s, "H", "X", 1))
	fmt.Println(s)
	fmt.Println(strings.Contains(s, "world"))
	fmt.Println("world\n" + "world")
	fmt.Println(`world
	world`)
	fmt.Println(`"`)

}
