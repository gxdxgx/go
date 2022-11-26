package main

import "fmt"

// 定数 型を宣言しない
const Pi = 3.14

const (
	Username = "test_user"
	Password = "test_pass"
)

func main() {
	fmt.Println(Pi, Username, Password)

}
