package mypackage

import (
	"fmt"
	"math"
)

// 小文字で定義すればこのファイル内のみで使用可能となる
type vertex3 struct {
	X, Y float64
}

// メソッドはレシーバ引数という特別な引数を関数に取る
func (v vertex3) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// 上記のメソッドを通常の関数として定義する場合はこうなる
// func Abs(v vertex3) float64 {
// 	return math.Sqrt(v.X*v.X + v.Y*v.Y)
// }

// structの型だけでなく、任意の型（type）にもメソッドを宣言できる
type myFloat float64

// レシーバを伴うメソッドの宣言は、レシーバ型が同じパッケージ内にある必要がある
func (f myFloat) Abs() float64 {
	if f < 0 {
		return float64(-f) // この記述法は型アサーション
	}
	return float64(f)
}

// ポインタレシーバでメソッドを宣言できる。ここではScaleというメソッド
// レシーバ自身を更新することが多いため、変数レシーバよりもポインタレシーバの方が一般的
// 変数レシーバでは元の変数のコピーを操作する。宣言した変数を変更するにはポインタレシーバにする必要がある
func (v *vertex3) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// ポインタレシーバを使う利点１：メソッドがレシーバの指す先の変数を変更する
// ポインタレシーバを使う利点２：メソッドの呼び出し毎に変数のコピーを避ける

// インターフェース型はメソッドの定義の集まり
// Abs()というメソッド名で返り値の型がfloat64であるメソッドが集まる
type Abser interface {
	Abs() float64
}

type I interface {
	M() // float64型とstring型を持つ構造体T型の返り値を持つMメソッドが定義される
	// interfaceには複数のメソッドを記述可能だが、型が存在しないとエラーになる
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return // 特に返り値のない場合はreturnで締める
	}
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func UseMethods() {
	fmt.Println("UseMethods")

	v := vertex3{3, 4}
	fmt.Println(v.Abs())

	f := myFloat(-math.Sqrt2)
	fmt.Println(f.Abs())

	v.Scale(10)
	fmt.Println(v.Abs())
	fmt.Println(v) // 定義した変数自体の値が更新されている！！

	p := &vertex3{3, 4} // ポインタ宣言
	p.Scale(3)
	fmt.Println(v, p) // ポインタレシーバの場合、変数またはポインタのレシーバとして取ることができる

	fmt.Println(p.Abs()) // 変数レシーバの場合でも同じく、変数またはポインタのレシーバとして取ることができる

	// aの値によって、それに対応するメソッドが実行される
	var a Abser
	a = f
	// ここではmyFloat型の値fを持ち、f.Abs()メソッドが実行される
	fmt.Println(a.Abs())

	a = p // (p.Scale(3)で値が更新された状態のもの)
	// ここでは*vertex3型のポインタが代入され、(*vertex3).Abs()メソッドが実行される
	fmt.Println(a.Abs())

	//
	var i I

	i = &T{"Hello"}
	i.M()

	i = F(math.Pi)
	i.M()

	var t *T // 構造体T型の変数、初期値はnil
	i = t    // インターフェースIにnilの構造体T型を代入
	i.M()    // nil値のためif文処理で<nil>と表示される。インターフェースの値自体はnilではない
	// nilインターフェースの値は、値も型も保持しない▶︎ランタイムエラーが発生する

	// ゼロ個のメソッドが指定されたインターフェース型は、空のインターフェースと呼ばれる
	var i2 interface{}
	describe(i2)

	// 空のインターフェースは、任意の型の値を保持できる
	// 未知の型の値を扱うコードで使用される。
	i2 = 42
	describe(i2)

	i2 = "hello"
	describe(i2)
}
