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

	// APIエンドポイント
	url := "https://jsonplaceholder.typicode.com/users/1"

	// GETリクエスト送信
	resp, err := http.Get(url)
	if err != nil {
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

	fmt.Printf("ID: %d\n", user.ID)
	fmt.Printf("Name: %s\n", user.Name)
	fmt.Printf("Email: %s\n", user.Email)
}
