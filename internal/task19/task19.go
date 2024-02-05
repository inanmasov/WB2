package task19

import "fmt"

// reverseString переворачивает строку с учетом символов Unicode
func reverseString(input string) string {
	// Преобразуем строку в слайс рун
	runes := []rune(input)

	// Переворачиваем слайс рун
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	// Возвращаем перевернутую строку
	return string(runes)
}

// Реализация программы, переворачивающей подаваемую на ход строку
func Task19() {
	// Пример использования
	inputString := "♨★✪❤✰✈"
	reversedString := reverseString(inputString)

	fmt.Printf("Исходная строка: %s\n", inputString)
	fmt.Printf("Перевернутая строка: %s\n", reversedString)
}
