package mypackage

import "fmt"

// 別ファイルで定義しているのでstruct名を変える必要がある
type Vertex2 struct {
	Lat, Long float64
}

var m map[string]Vertex2

// Golangにおけるマップの概念
func UseMaps() {
	fmt.Println("UseMaps")

	m = make(map[string]Vertex2)
	m["Bell Labs"] = Vertex2{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])

	var sampleMap = map[string]Vertex2{
		"Bell Labs": {
			40.68433, -74.39967,
		},
		"Google": {
			37.42202, -122.08408,
		},
	}
	fmt.Println(sampleMap)

	sampleMap2 := make(map[string]int)
	fmt.Println(sampleMap2)

	// マップへの値の挿入
	sampleMap2["Answer"] = 40
	fmt.Println(sampleMap2)

	// マップへの値の更新
	sampleMap2["Answer"] = 60
	fmt.Println(sampleMap2)

	// マップからの値の取得
	element := sampleMap2["Answer"]
	fmt.Println(element)

	// マップへの値の挿入（新しいkey）
	sampleMap2["Result"] = 99
	fmt.Println(sampleMap2)

	// マップのキー存在確認
	element2, exist := sampleMap2["Answer"] // 2つ目の返り値にbool値が返る
	fmt.Println(element2, exist)

	// マップの要素を削除
	delete(sampleMap2, "Answer")
	fmt.Println(sampleMap2)

	// マップのキー存在確認（ない場合）
	element2, exist = sampleMap2["Answer"] // 定義済みの変数なので代入式
	fmt.Println(element2, exist)           // 結果はint型のゼロ値「0」とfalse

	// マップのキー存在確認｜bool値のみ必要な場合は1つ目の返り値を捨てることができる
	_, exist2 := sampleMap2["Result"]
	fmt.Println(exist2)
}
