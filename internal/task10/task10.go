package task10

import (
	"fmt"
	"sort"
)

// Функция Task10 реализует объединение групп
func Task10() {
	// Создаем массив из чисел типа float64
	numbers := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	// map для хранения групп
	myMap := make(map[int][]float64)
	// Сортируем по возрастанию
	sort.Float64s(numbers)
	// Сохраняем максимально значение
	max := numbers[len(numbers)-1]

	// Проходим по всем отрезкам с заданным шагом
	for min, i := int(numbers[0]/10)*10, 0; ; min, i = min+10, i+1 {
		tmpArr, idx := separation(numbers, min)
		// Если в отрезке есть числа, сохраняем их в map
		if tmpArr != nil {
			myMap[min] = tmpArr
		}
		// Если рассматриваем последний отрезок
		if int(max)-min < 10 {
			break
		}
		// Если оставшиеся данные принадлежат одному отрезку
		if idx >= len(numbers) {
			break
		}
		// Удаляем ненужные данные
		numbers = numbers[idx:]
	}

	// Выводим map
	fmt.Println(myMap)
}

// Возвращаем слайс, входящий в группу с соответствующим шагом
func separation(nums []float64, border int) ([]float64, int) {
	var tmpArray []float64
	var i int

	if border >= 0 {
		border += 10
	}

	for i = 0; i < len(nums); i++ {
		if nums[i] > float64(border) {
			break
		}
		tmpArray = append(tmpArray, nums[i])
	}
	return tmpArray, i
}
