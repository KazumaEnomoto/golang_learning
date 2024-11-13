package mypackage

import "fmt"

// Golangにおけるポインタの概念
func UsePointer() {
	fmt.Println("UsePointer")

	a := 24

	p := &a         // ポインタ
	fmt.Println(*p) // ポインタを通して値を読み込む

	*p = 32 // ポインタを通して値を更新
	fmt.Println(a)
}
