package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func RequestSampleData() {
	fmt.Println("RequestSampleData")

	// APIエンドポイント
	url := "https://jsonplaceholder.typicode.com/users/1"

	// GETリクエスト送信
	resp, err := http.Get(url) // http.Getはレスポンスとエラーの2つの返り値を持つ
	if err != nil {            // errの値がnilの場合には正常
		fmt.Println("Error fetching data", err)
		return
	}
	defer resp.Body.Close()

	// ステータスコードチェック
	if resp.StatusCode != 200 {
		fmt.Println("Error: status code", resp.StatusCode)
		return
	}

	// レスポンスをデコード
	var user User
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}
	// 代入式ではerrにデコード結果のエラー情報が返され、デコードしたデータがuserに格納される
	// user構造体のポインタを指定すると、データ構造をコピーせずにデータ構造のメモリアドレスをコピーするだけで済むため、メモリ使用量と処理速度が向上する

	fmt.Printf("ID: %d\n", user.ID)
	fmt.Printf("Name: %s\n", user.Name)
	fmt.Printf("Email: %s\n", user.Email)
}
