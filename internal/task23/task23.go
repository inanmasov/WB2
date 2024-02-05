package task23

import "fmt"

func removeElement(arr []int, index int) []int {
	if index < 0 || index > len(arr) {
		fmt.Println("index out of range")
		return arr
	}
	return append(arr[:index], arr[index+1:]...)
}

// Реализация программы, удаляющей i-ый элемент из слайса
func Task23() {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	arr = removeElement(arr, 10)
	fmt.Println(arr)
}
