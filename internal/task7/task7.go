package task7

import (
	"fmt"
	"sync"
)

// Map представляет безопасный для конкурентного доступа map
type Map struct {
	mu sync.Mutex
	mp map[string]int
}

// NewMap конструктор экземпляра SafeMap
func NewMap() *Map {
	return &Map{
		mp: make(map[string]int),
	}
}

// Устанавливаем значение в map
func (m *Map) setMap(key string, val int) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.mp[key] = val
}

// Получаем значение из map
func (m *Map) GetMap(key string) (int, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	val, err := m.mp[key]
	return val, err
}

// Функция Task7 реализует безопасный конкурентный доступ map
func Task7() {
	mp := NewMap()

	// Создание WaitGroup для ожидания завершения горутин
	var wg sync.WaitGroup

	// Запуск горутин
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			// Параллельная запись в map без синхронизации
			mp.setMap(fmt.Sprintf("%d", index), index)
		}(i)
	}

	// Ожидание завершения всех горутин
	wg.Wait()

	// Вывод итоговой map
	fmt.Println("Итоговая map:", mp)
}
