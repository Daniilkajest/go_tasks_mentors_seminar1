package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Print("Введите целое число N: ")
	fmt.Scan(&n)

	// Проверка на знак
	if n > 0 {
		fmt.Println("Число положительное")
	} else if n < 0 {
		fmt.Println("Число отрицательное")
	} else {
		fmt.Println("Число равно нулю")
	}

	// Проверка на четность
	if n%2 == 0 {
		fmt.Println("Число чётное")
	} else {
		fmt.Println("Число нечётное")
	}

	// Проверка делимости на 3 и 5
	if n%3 == 0 && n%5 == 0 {
		fmt.Println("Число делится на 3 и на 5")
	} else {
		fmt.Println("Число НЕ делится одновременно на 3 и на 5")
	}
}