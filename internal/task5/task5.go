package task5

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Точка входа в Task5
func Task5() {
	var timeSec int
	fmt.Print("Укажите время работы в секундах: ")
	fmt.Scan(&timeSec)

	var wg sync.WaitGroup
	// Канал для передачи данных между воркерами и писателем
	ch := make(chan any)

	// Запуск горутины для писателя
	wg.Add(1)

	go func(out chan<- any, t int) {
		// Создаем новый таймер
		timer := time.NewTimer(time.Duration(t) * time.Second)
		for {
			select {
			case <-timer.C:
				wg.Done()
				close(ch)
				return
			default:
				out <- rand.Int()
			}

		}
	}(ch, timeSec)

	wg.Add(1)
	go func() {
		defer wg.Done()
		reader(ch)
	}()

	// Ожидание завершения всех горутин
	wg.Wait()
}

// Функция reader обрабатывает данные из канала и выводит их на экран
func reader(in <-chan any) {
	for val := range in {
		fmt.Printf("Вывод: %v\n", val)
	}
}
