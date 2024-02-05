package task26

import (
	"fmt"
	"sort"
	"strings"
)

// Сортируем входную строку и попарно проверяем все символы на идентичность
func Task26_2() {
	var text string
	fmt.Print("Введите строку: ")
	fmt.Scan(&text)

	fmt.Println(checkUniqueness2(text))
}

// Проверка на уникальность
func checkUniqueness2(text string) bool {
	// Приводим строку к нижнему регистру перед сортировкой
	sortedString := sortString(strings.ToLower(text))

	// Проверяем наличие повторяющихся символов после сортировки
	for i := 0; i < len(sortedString)-1; i++ {
		if sortedString[i] == sortedString[i+1] {
			return false
		}
	}

	return true
}

// Сортировка симовлов в строке
func sortString(s string) string {
	sChars := strings.Split(s, "")
	sort.Strings(sChars)
	return strings.Join(sChars, "")
}
