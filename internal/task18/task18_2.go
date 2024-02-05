package task18

import (
	"fmt"
	"sync"
)

// Реализация структуры-счетчика
func Task18_2() {
	// Создаем экземпляр счетчика
	counter := Counter1{}

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

// Counter - структура счетчика
type Counter1 struct {
	value int
	mu    sync.Mutex
}

// Increment - инкрементирует значение счетчика
func (c *Counter1) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

// Get - возвращает текущее значение счетчика
func (c *Counter1) Get() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}
