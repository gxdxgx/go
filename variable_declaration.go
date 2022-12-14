package main

import "fmt"

func foo() {
	// ショートの書き方、型が明確な時
	xi := 1
	xf64 := 1.2
	xs := "test"
	xt, xf := true, false
	fmt.Println(xi, xf64, xs, xt, xf)
}

func main() {
	// 通常の書き方、型を宣言したい時
	var (
		i    int     = 1
		f64  float64 = 1.2
		s    string  = "test"
		t, f bool    = true, false
	)
	i = 2
	fmt.Println(i, f64, s, t, f)
	foo()

	fmt.Printf("%T", f64)
}
