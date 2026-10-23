package main

import (
	"fmt"
)

func main() {
	var balance, amount float64
	var choice int

	fmt.Print("Введите ваш начальный баланс: ")
	fmt.Scan(&balance)

	for {
		fmt.Println("\nВыберите действие:")
		fmt.Println("1. Посмотреть баланс")
		fmt.Println("2. Пополнить баланс")
		fmt.Println("3. Снять средства")
		fmt.Println("4. Выход")
		fmt.Print("Ваш выбор: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Printf("Текущий баланс: %.2f\n", balance)
		case 2:
			fmt.Print("Введите сумму для пополнения: ")
			fmt.Scan(&amount)
			if amount > 0 {
				balance += amount
				fmt.Println("Баланс успешно пополнен.")
			} else {
				fmt.Println("Сумма должна быть положительной.")
			}
		case 3:
			fmt.Print("Введите сумму для снятия: ")
			fmt.Scan(&amount)
			if amount <= 0 {
				fmt.Println("Сумма должна быть положительной.")
			} else if amount > balance {
				fmt.Println("Ошибка: недостаточно средств на балансе.")
			} else {
				balance -= amount
				fmt.Println("Средства успешно сняты.")
			}
		case 4:
			fmt.Println("До свидания!")
			return
		default:
			fmt.Println("Неизвестное действие. Попробуйте снова.")
		}
	}
}