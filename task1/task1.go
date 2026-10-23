package main

import (
	"fmt"
)

func main() {
	var a, b float64
	var op string

	fmt.Println("Введите первое число, операцию (+, -, *, /) и второе число через пробел:")
	fmt.Scan(&a, &op, &b)

	switch op {
	case "+":
		fmt.Printf("Результат: %.2f\n", a+b)
	case "-":
		fmt.Printf("Результат: %.2f\n", a-b)
	case "*":
		fmt.Printf("Результат: %.2f\n", a*b)
	case "/":
		if b == 0 {
			fmt.Println("Ошибка: деление на ноль!")
		} else {
			fmt.Printf("Результат: %.2f\n", a/b)
		}
	default:
		fmt.Println("Ошибка: неизвестная операция")
	}
}