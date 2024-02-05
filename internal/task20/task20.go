package task20

import (
	"fmt"
	"strings"
)

// reverseString переворачивает слова с учетом символов Unicode
func reverseString(input string) string {

	words := strings.Split(input, " ")

	// Создание двумерного массива рун
	runes := make([][]rune, len(words))

	for i, word := range words {
		runes[i] = []rune(word)
	}

	// Переворачиваем слайс рун
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	var str string
	for _, val := range runes {
		str += string(val) + " "
	}

	// Возвращаем перевернутую строку
	return str[:len(str)-1]
}

// Реализация программы, которая переворачивает слова в строке
func Task20() {
	// Пример использования
	inputString := "snow dog sun"
	reversedString := reverseString(inputString)

	fmt.Printf("Исходная строка: %s\n", inputString)
	fmt.Printf("Перевернутая строка: %s\n", reversedString)
}
