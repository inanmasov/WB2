package task16

import "fmt"

// Функция Task16 демонстрирует использование алгоритма быстрой сортировки (QuickSort).
func Task16() {
	arr := []int{2, 9, 25, 1, 5, 4, 3, 2, 9, 11, 15, 42, 5, 6, 25}

	quickSort(arr, 0, len(arr)-1)

	fmt.Println(arr)
}

// Функция quickSort выполняет рекурсивную сортировку массива методом QuickSort
func quickSort(arr []int, low int, high int) {
	if low < high {
		q := partition(arr, low, high)
		quickSort(arr, low, q)
		quickSort(arr, q+1, high)
	}
}

// Функция partition выполняет разделение массива для алгоритма QuickSort
func partition(arr []int, low int, high int) int {
	mid := arr[(low+high)/2]
	for low <= high {
		for ; arr[low] < mid; low++ {

		}
		for ; arr[high] > mid; high-- {

		}
		if low >= high {
			break
		}
		arr[low], arr[high] = arr[high], arr[low]
		low++
		high--
	}
	return high
}
