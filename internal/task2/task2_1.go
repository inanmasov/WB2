package task2

import (
	"fmt"
	"sync"
	"time"
)

// Точка входа в Task2
func Task2_1() {
	// Засекаем время начала выполнения участка кода
	startTime := time.Now()

	numbers := []int{2, 4, 6, 8, 10}

	var wg sync.WaitGroup

	// Проходимя по всем элементам массива
	for _, v := range numbers {
		// Увелчиваем счетчик горутин
		wg.Add(1)
		go func(v int) {
			// По окончании функции уменьшаем счетчик  горутин
			defer wg.Done()
			fmt.Println(v * v)
		}(v)
	}

	// Ожидание завершения всех горутин.
	wg.Wait()

	// Засекаем время завершения выполнения участка кода
	endTime := time.Now()
	// Выводим затраченное время
	fmt.Printf("Time taken: %s\n", endTime.Sub(startTime))
}
