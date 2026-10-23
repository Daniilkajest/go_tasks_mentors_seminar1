package main

import (
	"fmt"
)

func main() {
	var n, val, posCount, negCount, zeroCount int

	fmt.Print("Введите количество чисел N: ")
	fmt.Scan(&n)

	fmt.Println("Введите числа:")
	for i := 0; i < n; i++ {
		fmt.Scan(&val)
		if val > 0 {
			posCount++
		} else if val < 0 {
			negCount++
		} else {
			zeroCount++
		}
	}

	fmt.Printf("Положительных: %d\n", posCount)
	fmt.Printf("Отрицательных: %d\n", negCount)
	fmt.Printf("Нулей: %d\n", zeroCount)
}