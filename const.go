package main

import "fmt"

func main() {
	// 상수
	// const는 사용 초기화, 한 번 선언 후 값 변경 금지. 고정된 값 관리.
	const a string = "Test" // 바로 초기화가되어야 한다.
	const b = "Test2"
	const c int32 = 10 * 10
	// 함수 리턴받을 수 없다.
	const e = 3.14
	const f = false

	/*
		에러 발생 => 상수는 선언과 동시에 초기화 되어야 한다.
		const g string
		g = "Test3"
	*/

	fmt.Println("a : ", a, "b : ", b, "c : ", c, "e : ", e, "f : ", f)

	// 상수 2
	const A, B, C int = 10, 20, 30
	const D, E = true, "abc"
	const (
		x, y int16 = 50, 90
		i, k       = "Data", 888
	)
	fmt.Println("A: ", A, "B: ", B, "C: ", C, "D: ", D, "E: ", E, "x: ", x, "y: ", y, "i : ", i, "k : ", k)
}
