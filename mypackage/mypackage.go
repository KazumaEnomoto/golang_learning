package mypackage

import "fmt"

// main以外のpackageを使用するにはディレクトリを作成してその配下で作成する

// 大文字から始めるものは外部からアクセス可能
func MyPublicFunction() {
	fmt.Println("PublicFunction")
}

// 小文字から始めるものは、それを定義したパッケージ内のみで使用可能
func myPrivateFunction() {
	fmt.Println("PrivateFunction")
}
