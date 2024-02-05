package task17

import (
	"fmt"
	"sort"
)

// Функция Task16 демонстрирует использование бинарного поиска
func Task17() {
	arr := []int{2, 9, 25, 1, 5, 4, 3, 11, 15, 42, 6}
	sort.Ints(arr)
	fmt.Println(arr)

	for _, v := range arr {
		fmt.Printf("Элемент %d имеет индекс: %d\n", v, binarySearch(arr, v))
	}

	fmt.Printf("Элемента %d нет в массиве: %d\n", 0, binarySearch(arr, 0))
}

// Бинарный поиск
func binarySearch(arr []int, value int) int {
	// Границы
	low, high := 0, len(arr)-1

	for low <= high {
		// Середина
		mid := (low + high) / 2

		// Элемент найден
		if arr[mid] == value {
			return mid
		} else if arr[mid] < value {
			// Искомый элемент находится в правой половине
			low = mid + 1
		} else {
			// Искомый элемент находится в левой половине
			high = mid - 1
		}
	}

	return -1
}
