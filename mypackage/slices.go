package mypackage

import "fmt"

// GolangにおけるSlice（スライス）の概念

// 配列（Array）
func UseArray() {
	fmt.Println("UseArray")

	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}

// スライス（Slice）
func UseSlice() {
	fmt.Println("UseSlice")

	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4] // 要素の内1から3を含むスライスが作られる
	fmt.Println(s)            // [3, 5, 7]

	fruits := [4]string{
		"apple",
		"banana",
		"orange",
		"grape",
	}
	fmt.Println(fruits)

	a := fruits[0:2]
	b := fruits[1:3]
	fmt.Println(a, b)

	// スライスの要素を変更すると、元の配列の要素も変更される（他のスライスにも影響が出る）
	b[0] = "XXXX"
	fmt.Println(a, b)
	fmt.Println(fruits)

	c := primes[:] // スライスは上限・下限を省略できる（上限はスライスの長さ、下限は0）
	fmt.Println(c)

}
