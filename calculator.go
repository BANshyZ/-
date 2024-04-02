package main

import (
	"fmt"
)

func main() {
	var a, b float64
	var result float64
	var op string
	fmt.Print("Введите первое число: ")
	fmt.Scanln(&a)

	fmt.Print("Введите второе число: ")
	fmt.Scanln(&b)

	fmt.Print("Какое действие выполнить?:\n1 - сложение (+)\n2 - вычитание (-)\n3 - умножение (*)\n4 - деление (/)\n")
	fmt.Scanln(&op)
	switch op {
	case "1":
		result = a + b
		fmt.Println("Результат: ", a, "+", b, "=", result)
	case "2":
		result = a - b
		fmt.Println("Результат: ", a, "-", b, "=", result)
	case "3":
		result = a * b
		fmt.Println("Результат: ", a, "*", b, "=", result)
	case "4":
		if b != 0 {
			result = a / b
			fmt.Println("Результат: ", a, "/", b, "=", result)
		} else {
			fmt.Println("На ноль не делится")
			return
		}
	default:
		fmt.Println("Неправильная операция")
		return
	}
}
