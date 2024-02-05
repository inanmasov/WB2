package task26

import (
	"fmt"
	"strings"
)

// Проверка на уникальность
func checkUniqueness1(text string) bool {
	// Приводим все символы к одному регистру
	text = strings.ToLower(text)
	// Создаем map для хранения символов
	myMap := make(map[rune]int)

	// Сохраняем символы в map
	for _, value := range text {

		myMap[value]++
	}
	// Проверям встретился ли какой-нибудь символ больше одного раза
	for _, value := range myMap {
		if value > 1 {
			return false
		}
	}

	return true
}

// Реализация программы, которая проверяет, что все символы в строке уникальные
func Task26_1() {
	var text string
	fmt.Print("Введите строку: ")
	fmt.Scan(&text)

	fmt.Println(checkUniqueness1(text))
}
