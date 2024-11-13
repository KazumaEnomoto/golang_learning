package main

// packageのインポート
// 自モジュール内のパッケージをインポートする場合はディレクトリ名で指定する
import (
	"fmt"
	"golang_learning/mathfunction"
	"golang_learning/mypackage"
	"runtime"
	"time"
)

func main() {
	fmt.Printf("Hello World\n")

	// Golangにおける変数の定義
	var message string = "I'm learning Golang."
	fmt.Println(message)

	var number int = 1234
	fmt.Println(number)

	var isActive bool = false
	fmt.Println(isActive)

	// 下記のように複数の変数を定義することもできる
	var i, j int = 1, 2
	fmt.Println(i, j)

	// 複数定義の省略形
	foo, bar := "foo", "bar"
	fmt.Println(foo, bar)

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

	// Golangにおける定数の定義（定数は:=を使って宣言できない）
	const id = 1                       // 型宣言しない場合
	const password string = "password" // 型宣言する場合
	fmt.Println(id)
	fmt.Println(password)

	// 現在時刻の取得
	now := time.Now()
	fmt.Println(now)

	// 自モジュール内のパッケージ内関数の使用
	mypackage.MyPublicFunction()

	a := 9
	b := mathfunction.Cubed(a)
	fmt.Println(b)

	c := 12
	d := 34
	e := mathfunction.Add(c, d)
	fmt.Println(e)

	mypackage.UsePrivateFunction()

	// 型と中身を見る方法
	v := 42
	w := "message"
	fmt.Printf("%T\n", v)
	fmt.Printf("%v\n", v)
	fmt.Printf("%v\n", w)

	// Golangにおけるforループ処理
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	// 初期化と後処理は省略可能（つまりwhile）
	sum2 := 1
	for sum2 < 10 {
		sum2 += sum2
	}
	fmt.Println(sum2)

	// Golangにおけるif条件分岐処理
	number2 := 0
	if number2 > 0 {
		fmt.Println("正の数")
	} else if number2 == 0 {
		fmt.Println("ゼロ")
	} else {
		fmt.Println("負の数")
	}

	// Golangではif文の前に代入文を書くことができる（if文スコープ内で使用可能）
	number3 := 31
	if v := number3 % 3; v == 0 {
		fmt.Printf("余りが%vのため3の倍数\n", v)
	} else {
		fmt.Printf("余りが%vのため3の倍数ではない\n", v)
	}

	// Golangにおけるswitch文（switch文でも代入文を書くことができる）
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.\n", os)
	}

	// defer｜渡した関数の実行を呼び出し元の関数の終わりまで遅延させる
	delayPrint()

	// deferへ渡した関数はスタックされて、最後にスタックしたものから順番に実行される
	delayForLoop()

	// Golangにおけるポインタの概念
	mypackage.UsePointer()

	// Golangにおけるstructの概念
	mypackage.UseStruct()

	// Golangにおけるarray（配列）の使い方
	mypackage.UseArray()

	// Golangにおけるsliceの概念
	mypackage.UseSlice()
}

func delayPrint() {
	defer fmt.Println("world")
	fmt.Println("hello")
}

func delayForLoop() {
	fmt.Println("start for loop")

	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
