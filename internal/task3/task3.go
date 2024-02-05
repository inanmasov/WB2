package task3

import (
	"fmt"
	"sync"
)

// Точка входа в Task3
func Task3() {
	numbers := []int{2, 4, 6, 8, 10}

	var wg sync.WaitGroup
	ch := make(chan int, len(numbers))

	// Проходимя по всем элементам массива
	for _, v := range numbers {
		// Увелчиваем счетчик горутин
		wg.Add(1)
		go func(v int, ch chan<- int) {
			// По окончании функции уменьшаем счетчик  горутин
			defer wg.Done()
			// Записываем в канал квадрат числа
			ch <- v * v
		}(v, ch)
	}

	// Горутина, закрывающая канал после завершения всех горутин
	go func() {
		wg.Wait()
		close(ch)
	}()

	// Читаем результаты из канала и выводим их
	num := 0
	for result := range ch {
		num += result
	}
	fmt.Println(num)
}
