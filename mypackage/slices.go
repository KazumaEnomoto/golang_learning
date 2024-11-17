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

	// スライスは上限・下限を省略できる（上限はスライスの長さ、下限は0）
	c := primes[:]
	fmt.Println(c)

	// スライスは長さ（length）と容量（capacity）を持ち、len, capで得ることができる
	slicePrimes := []int{2, 3, 5, 7, 11, 13}
	PrintSlice(slicePrimes)

	// スライスの長さ0にする、容量は変わらない
	slicePrimes = slicePrimes[:0]
	PrintSlice(slicePrimes)

	// スライスの長さを4にする、容量は変わらない
	slicePrimes = slicePrimes[:4]
	PrintSlice(slicePrimes)

	// スライスの長さを8にする▶︎▶︎▶︎容量が6のためエラーとなる
	// slicePrimes = slicePrimes[:8]

	// スライスの長さを6（上限）にする、容量は変わらない（定義時のスライス）
	slicePrimes = slicePrimes[:6]
	PrintSlice(slicePrimes)

	// スライスの1つ目の値を取り除く、長さ・容量はともに5になる
	slicePrimes = slicePrimes[1:]
	PrintSlice(slicePrimes)
}

func PrintSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
