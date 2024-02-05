package task2

import (
	"fmt"
	"sync"
	"time"
)

func calculateSquare(number int, wg *sync.WaitGroup, ch chan<- int) {
	// Уменьшаем счетчик ожидания после завершения горутины
	defer wg.Done()
	square := number * number
	// Отправляем результат в канал
	ch <- square
}

func Task2_2() {
	// Засекаем время начала выполнения участка кода
	startTime := time.Now()

	numbers := []int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup
	ch := make(chan int, len(numbers))

	// Запускаем горутины для вычисления квадратов чисел
	for _, number := range numbers {
		wg.Add(1)
		go calculateSquare(number, &wg, ch)
	}

	// Запускаем горутину для закрытия канала после завершения всех горутин
	go func() {
		// Ждем завершения всех горутин
		wg.Wait()
		// Закрываем канал после завершения всех горутин
		close(ch)
	}()

	// Читаем результаты из канала и выводим их
	for result := range ch {
		fmt.Printf("Square: %d\n", result)
	}

	// Засекаем время завершения выполнения участка кода
	endTime := time.Now()
	// Выводим затраченное время
	fmt.Printf("Time taken: %s\n", endTime.Sub(startTime))
}
