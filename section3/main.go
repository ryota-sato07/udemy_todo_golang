package main

import (
	"fmt"
	"time"
)

/**
 * 変数
 */
// 関数内でしか暗黙定義は利用できない
// i5 := 500

var i5 int = 500

/**
 * outer関数
 */
func outer() {
	var s4 string = "outer"
	fmt.Println(s4)
}

/**
 * main関数
 */
func main() {

	//-- [ variables ] ---------------------------------------------------------
	/**
	 * 明示的な変数定義
	 */
	var i int = 100
	fmt.Println(i)

	var s string = "Hello Go"
	fmt.Println(s)

	// tuple のように複数の変数を定義できる
	var t, f bool = true, false
	fmt.Println(t, f)

	var (
		i2 int    = 200
		s2 string = "Golang"
	)
	fmt.Println(i2, s2)

	var i3 int
	var s3 string
	fmt.Println(i3, s3)

	i3 = 300
	s3 = "Go"
	fmt.Println(i3, s3)

	//-- [ variables ] ---------------------------------------------------------
	/**
	 * 暗黙的な変数定義
	 */
	i4 := 400
	fmt.Println(i4)

	i4 = 450
	fmt.Println(i4)

	// 暗黙を再定義することはできない
	// i4 := 500
	// fmt.Println(i4)

	// int型で型推論されたものは再代入できない
	i4 = 450
	fmt.Println(i4)

	fmt.Println(i5)

	outer()

	fmt.Println(time.Now())
}
