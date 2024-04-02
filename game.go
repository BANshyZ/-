package main

import (
	"fmt"
)

var charName string
var doorOpen bool

func main() {
	doorOpen = false
	var act string
	fmt.Println("Вы проснулись в неизвестнной комнате...\nВы пытаетесь вспомнить своё имя...\nВаше имя - ")
	fmt.Scan(&charName)
	fmt.Println("Отлично, " + charName + "\nпохоже вы вспомнили как вас зовут\nТеперь вам нужно выбраться из этой комнаты, вы осмотрелись, теперь можно")
	for doorOpen == false {
		fmt.Println("a. Открыть дверь\nb. Заглянуть под кровать\nc. Открыть ящик в углу комнаты\nd. Открыть вентиляцию\ne. Взглянуть на тумбочку\nf. Взглянуть на статую рядом с дверью\nЧто сделать?")
		fmt.Scanln(&act)
		switch act {
		case "a":
			openDoor()
		case "b":
			checkBed()
		case "c":
			openBox()
		case "d":
			openVent()
		case "e":
			checkTable()
		case "f":
			checkStatue()
		}
	}
}

func openDoor() {

	if doorKey == 1 {
		doorOpen = true
		fmt.Println("Вы использовали ключ\nДверь открылась, " + charName + ", теперь вы свободны!!!")
	} else {
		fmt.Println("Чтобы открыть дверь, нужно найти, чем её открыть")
	}
}

var art1 int

func checkBed() {
	fmt.Println(charName + ", вы нашли первый артефакт под кроватью")
	art1++
}

var doorKey int

func openBox() {
	if boxKey == 1 {
		doorKey++
		fmt.Println(charName + ", вы открыли ящик, в нём ключ от двери")
	} else {
		fmt.Println("Нужно найти ключ для этого ящика")
	}
}

var count int
var art2 int

func openVent() {
	count++
	if count == 3 {
		fmt.Println(charName + ", вентиляция поддалась, вы нашли второй артефакт")
		art2++
	} else {
		fmt.Println("Вы пытаетесь открыть вентиляцию\nОткрыть вентиляцию не получается\nНужно попробовать её расшатать")
	}
}

var art3 int

func checkTable() {
	art3++
	fmt.Println(charName + ", вы нашли третий артефакт на тумбочке.")
}

var boxKey int

func checkStatue() {
	fmt.Println("Вы пытаетесь активировать статую...")
	if art1 == 1 && art2 == 1 && art3 == 1 {
		boxKey++
		fmt.Println(charName + ", вы активировали статую тремя артефактами и получили ключ от ящика.")
	} else {
		fmt.Println("В статуе три выемки для артефактов, нужно их найти, чтобы что-то получилось")
	}
}
