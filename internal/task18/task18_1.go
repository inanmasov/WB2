package task18

import (
	"fmt"
	"sync"
)

// Counter - структура счетчика
type Counter struct {
	Inc func() int
	Get func() int
}

// создаем замыкание
func NewCounter() Counter {
	count := 0
	var mu sync.Mutex
	incFunc := func() int {
		mu.Lock()
		count++
		mu.Unlock()
		return count
	}
	getFunc := func() int {
		return count
	}

	return Counter{Inc: incFunc, Get: getFunc}
}

// Реализация структуры-счетчика
func Task18_1() {
	// Создаем экземпляр счетчика
	counter := NewCounter()

	var wg sync.WaitGroup

	// Запускаем 100 горутин для инкрементирования счетчика
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Inc()
		}()
	}

	// Ожидаем завершения всех горутин
	wg.Wait()
	// Выводим значение счетчика
	fmt.Println(counter.Get())
}
