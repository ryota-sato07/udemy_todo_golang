package main

import "fmt"

/**
 * main関数
 */
func main() {

	var i int = 100
	fmt.Println(i + 50)

	var i2 int64 = 100
	fmt.Println(i2)

	/**
	 * int型だとしても、環境依存の int型 と 明示的に指定した int型とでは、別の型と判断されるためエラーとなる
	 */
	// fmt.Println(i + i2)

	fmt.Printf("%T\n", i2)
}
