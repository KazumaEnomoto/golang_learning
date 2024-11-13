package mypackage

import (
	"fmt"
)

// Golangにおけるstruct（構造体）の概念

// struct（構造体）はfield（フィールド）の集まり
type Vertex struct {
	X int
	Y int
}

func UseStruct() {
	v := Vertex{1, 2}
	fmt.Println(v)

	v.X = 4 // fieldにアクセスするには「.（ドット）」を用いる
	fmt.Println(v)

	p := &v
	p.X = 12 // ポインタを通してアクセス可能｜(*p).Xで表記できるがp.Xと省略することができる
	fmt.Println(v)

	var (
		v1 = Vertex{1, 2}
		v2 = Vertex{X: 2} // Name:構文を使用して一部のfieldを初期化
		v3 = Vertex{}     // 初期値がないためint型のゼロ値が初期値となる
	)

	fmt.Println(v1, v2, v3)
}
