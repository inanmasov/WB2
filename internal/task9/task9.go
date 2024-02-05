package task9

import (
	"fmt"
	"sync"
)

// Функция Task9 реализует конвейер чисел
func Task9() {
	// Переменная для хранения количества воркеров
	var countNum int
	fmt.Print("Укажите количество чисел в массиве: ")
	fmt.Scan(&countNum)

	array := make([]int, countNum)

	fmt.Printf("Введите %d целых чисел:\n", countNum)

	// Вводим значения с клавиатуры и записываем их в массив
	for i := 0; i < countNum; i++ {
		fmt.Printf("Элемент %d: ", i+1)
		_, err := fmt.Scan(&array[i])
		if err != nil {
			fmt.Println("Ошибка ввода:", err)
			return
		}
	}

	ch1 := make(chan int)
	ch2 := make(chan int)

	// Создание WaitGroup для ожидания завершения горутин
	var wg sync.WaitGroup

	// Записываем данные из массива в первый канал
	wg.Add(1)
	go func(ch1 chan<- int, array []int) {
		defer wg.Done()
		for _, val := range array {
			ch1 <- val
		}
		// Закрываем канал ch1 после завершения записи
		close(ch1)
	}(ch1, array)

	// Берем данные из первого канала и записываем их квадраты во второй
	wg.Add(1)
	go func(ch1 <-chan int, ch2 chan<- int) {
		defer wg.Done()
		for val := range ch1 {
			ch2 <- val * val
		}
		// Закрываем канал ch2 после завершения записи
		close(ch2)

	}(ch1, ch2)

	// Выводим данные из второго канала
	wg.Add(1)
	go func(ch2 <-chan int) {
		defer wg.Done()
		for val := range ch2 {
			fmt.Println(val)
		}
	}(ch2)

	// Ожидание завершения всех горутин
	wg.Wait()
}
