package main

import (
	"fmt"
)

func main() {
	var n, val, max, min int

	fmt.Print("Введите количество чисел N: ")
	fmt.Scan(&n)

	if n <= 0 {
		fmt.Println("Количество чисел должно быть больше нуля")
		return
	}

	fmt.Println("Введите числа:")
	for i := 0; i < n; i++ {
		fmt.Scan(&val)
		if i == 0 {
			max = val
			min = val
		} else {
			if val > max {
				max = val
			}
			if val < min {
				min = val
			}
		}
	}

	fmt.Printf("Максимальное значение: %d\n", max)
	fmt.Printf("Минимальное значение: %d\n", min)
}