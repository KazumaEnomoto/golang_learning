package main

import "fmt"

func main() {
	fmt.Printf("Hello World\n")

	// Golangにおける変数の定義
	var message string = "I'm learning Golang."
	fmt.Println(message)

	var number int = 1234
	fmt.Println(number)

	var isActive bool = false
	fmt.Println(isActive)

	// 下記のように「var」を省略して記述できる（Goの書き方）
	// ただし、関数内でのみ使用できるため、パッケージレベルの変数を定義する場合は「var」を使用する
	x := 333
	fmt.Println(x)

	// Golangにはゼロ値の概念があり、変数に初期値を割り当てないと、その型のゼロ値が自動的に割り当てられる
	var i_number int     // ゼロ値：0
	var f_number float64 // ゼロ値：0.0
	var text string      // ゼロ値：""（空文字）
	var boolean bool     // ゼロ値：false
	fmt.Println(i_number)
	fmt.Println(f_number)
	fmt.Println(text)
	fmt.Println(boolean)

	// Golangにおける定数の定義
	const id = 1                       // 型宣言しない場合
	const password string = "password" // 型宣言する場合
	fmt.Println(id)
	fmt.Println(password)
}
