package mypackage

import (
	"fmt"
	"math/rand"
)

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

	// スライスの長さを6にする▶︎▶︎▶︎容量が5に減っているためエラーとなる
	// slicePrimes = slicePrimes[:6]

	// スライスの長さを5（上限）にする、容量は変わらない（定義時の1つ目の値が取り除かれている）
	slicePrimes = slicePrimes[:5]
	PrintSlice(slicePrimes)

	// スライスのゼロ値はnilになる
	var zeroSlice []int
	if zeroSlice == nil {
		fmt.Println("nil")
	}

	// ランダムなスライスを表示
	CreateRandomSlice()

	// スライスへの値の追加
	var slice1 []int
	PrintSlice(slice1)

	slice1 = append(slice1, 0)
	PrintSlice(slice1)

	slice1 = append(slice1, 1)
	PrintSlice(slice1)

	slice1 = append(slice1, 1, 2, 3)
	PrintSlice(slice1)

	// forループで使用するrange、マップ(map)をひとつずつ反復処理する（foreachみたいなもの）
	pow := []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow { // 1つ目はインデックス、2つ目はその要素
		fmt.Printf("2の%d乗は%d\n", i, v)
	}

	// rangeにおけるインデックスや値は「_」アンダーバーで捨てることができる
	slice2 := make([]int, 5)
	for i := range slice2 { // 2つ目（値）がいらない場合は省略可能
		slice2[i] = 2 * i
	}
	PrintSlice(slice2)
	for _, v := range slice2 { // インデックスがいらない場合は「_」で表記する
		fmt.Println(v)
	}
}

func PrintSlice(s []int) {
	fmt.Println("PrintSlice")
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

// ランダムな数字を取得して、その長さのスライスを取得する関数
func CreateRandomSlice() {
	fmt.Println("CreateRandomSlice")
	randomNumber := rand.Intn(10)
	fmt.Println(randomNumber)

	// スライスの長さを指定
	s := make([]int, randomNumber) // 値はint型のゼロ値「0」となる
	PrintSlice(s)

	// スライスの長さと容量を指定
	sliceLen := 5
	sliceCap := 5
	s2 := make([]int, sliceLen, sliceCap)
	PrintSlice(s2)
}
