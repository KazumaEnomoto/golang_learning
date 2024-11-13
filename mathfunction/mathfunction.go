package mathfunction

func Cubed(number int) int {
	result := number * number * number
	return result
}

func Add(x int, y int) int { // 引数の型が同じ場合には(x, y int)のように省略可能
	result := x + y
	return result
}
