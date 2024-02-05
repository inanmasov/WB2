package task11

import "fmt"

// Функция Task10 реализует пересечение двух неупорядоченных множеств
func Task11() {
	// Исходные множества.
	arr1 := []int{1, 7, 4, 12, 3, 2, 6, 23, 42, 25}
	arr2 := []int{2, 9, 1, 5, 4, 3, 11, 15, 42, 6, 25}

	var finalArray []int

	myMap := make(map[int]int)

	for _, val := range arr1 {
		myMap[val] = 1
	}

	for _, val := range arr2 {

		if _, ok := myMap[val]; ok {
			finalArray = append(finalArray, val)
		}
	}

	fmt.Println(finalArray)
}
