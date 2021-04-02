// package main

// func main() {
// 	1. 변수 초기화 방법. 초기화만 하거나 선언만 한 변수는 에러처리.

// 	var a int
// 	var b string
// 	a = 3
// 	b = "324"
// 	fmt.Println("hello")
// 	fmt.Println(a, b)
// 	fmt.Println(b)

// 	var c, d, e int = 1, 2, 3
// 	var k float32 = 11.4
// 	var J string = "Hi! Golang"
// 	fmt.Println(c, d, e)
// 	fmt.Println(k)
// 	fmt.Println(J)

// 	var PI = 3.14
// 	fmt.Println(PI)

// 	var isTrue = true
// 	fmt.Println(isTrue)
// 	////////////////////////////////////////////////////////////////////////////////////////////////
// 	2. 여러 개 선언
// 	var (
// 		name      string = "machine"
// 		height    int32
// 		weight    float32
// 		isRunning bool
// 	)

// 	height = 250
// 	weight = 350.23
// 	isRunning = true

// 	fmt.Println("name: ", name, "height: ", height, "weight: ", weight, "isRunning: ", isRunning)
// 	////////////////////////////////////////////////////////////////////////////////////////////////
// 	3. 짧은 선언 (중요!!!)
// 	전역으로 사용 불가하고, 함수 안에서만 사용. 선언 후 재선언시 할당 예외 발생함.
// 	주로 제한된 범위의 함수내에서 사용할 경우 코드 가독성을 높일 수 있다. -> 추천.

// 	shortVar1 := 3
// 	shortVar2 := "Test"
// 	shortVar3 := true

// 	// shortVar1 := 10 -> 짧은선언 한 변수에 대해 재할당 할 경우 예외 발생.

// 	fmt.Println("shortVar1: ", shortVar1, "shortVar2: ", shortVar2, "shortVar3: ", shortVar3)

// }
