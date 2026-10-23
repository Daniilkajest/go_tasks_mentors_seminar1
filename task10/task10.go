package main

import (
	"fmt"
)

func main() {
	var n, val, sum, max, min, evenCount, oddCount int

	fmt.Print("Введите количество чисел N: ")
	fmt.Scan(&n)

	if n <= 0 {
		fmt.Println("Количество чисел должно быть больше нуля")
		return
	}

	fmt.Println("Введите числа:")
	for i := 0; i < n; i++ {
		fmt.Scan(&val)

		sum += val

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

		if val%2 == 0 {
			evenCount++
		} else {
			oddCount++
		}
	}

	avg := float64(sum) / float64(n)

	fmt.Printf("\n--- Статистика ---\n")
	fmt.Printf("Сумма: %d\n", sum)
	fmt.Printf("Среднее значение: %.2f\n", avg)
	fmt.Printf("Максимальное значение: %d\n", max)
	fmt.Printf("Минимальное значение: %d\n", min)
	fmt.Printf("Количество чётных чисел: %d\n", evenCount)
	fmt.Printf("Количество нечётных чисел: %d\n", oddCount)
}