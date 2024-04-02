package main

import (
	"fmt"
)

func main() {
	var a, b, result int
	var op string

	fmt.Print("Введите минимальное число: ")
	fmt.Scanln(&a)

	fmt.Print("Введите максимальное число: ")
	fmt.Scanln(&b)

	if a >= b {
		fmt.Println("Минимальное числло не может быть больше максимального")
	} else {
		fmt.Println("Введите действие:\n1. Вывод суммы всех целых чисел от ", a, " до ", b, ".\n2. Вывод суммы всех целых четных чисел в промежутке от ", a, " до ", b, ".\n3. Вывод сумму всех целых нечетных чисел в промежутке от ", a, " до ", b, ".")
		fmt.Scanln(&op)
		switch op {
		case "1":
			result = (a + b) * (b - a + 1) / 2
			fmt.Print("Суммы всех целых чисел от ", a, " до ", b, " : ", result)
		case "2":
			for a < b {
				if a%2 == 0 {
					result += a
				}
				a++
			}
			a -= b - 1
			fmt.Print("Суммы всех целых чётных чисел от ", a, " до ", b, " : ", result)
		case "3":
			for a < b {
				if a%2 != 0 {
					result += a
				}
				a++
			}
			a -= b - 1
			fmt.Print("Суммы всех целых нечётных чисел от ", a, " до ", b, " : ", result)
		}
	}
}
