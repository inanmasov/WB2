package task12

import "fmt"

// Функция Task12 реализует собственное множество
func Task12() {
	// Создание нового множества.
	set := NewPlenty()

	// Добавление элементов в множество.
	set.Add("cat", "cat", "dog", "cat", "tree")

	// Удаление элемента из множества.
	set.Delete("tree")

	// Проверка наличия элемента в множестве.
	fmt.Println(set.Contain("dog"))

	// Получение всех ключей множества.
	fmt.Println(set.GetKeys())

	// Вывод содержимого множества.
	fmt.Println(set.data)
}

// Тип Set представляет собой простую реализацию множества.
type Plenty struct {
	data map[any]int
}

// NewSet создает новый экземпляр множества.
func NewPlenty() *Plenty {
	return &Plenty{data: map[any]int{}}
}

// Добавляет элементы в множество.
func (s *Plenty) Add(values ...any) {
	for _, value := range values {
		s.data[value]++
	}
}

// Удаляет элементы из множества.
func (s *Plenty) Delete(keys ...any) {
	for _, key := range keys {
		if s.data[key] > 1 {
			s.data[key]--
		} else {
			delete(s.data, key)
		}
	}
}

// Проверяет наличие элемента в множестве.
func (s Plenty) Contain(data any) bool {
	_, ok := s.data[data]
	return ok
}

// Возвращает все ключи множества.
func (s *Plenty) GetKeys() []any {
	keys := make([]any, 0, len(s.data))

	for k, v := range s.data {
		for i := 0; i < v; i++ {
			keys = append(keys, k)
		}
	}

	return keys
}
