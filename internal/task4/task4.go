package task4

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

// Точка входа в Task4
func Task4() {
	// Создаем канал для получения сигналов
	sigChan := make(chan os.Signal, 1)

	// Регистрируем канал для сигналов завершения программы
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	// Переменная для хранения количества воркеров
	var workersCount int
	fmt.Print("Укажите количество воркеров: ")
	fmt.Scan(&workersCount)

	var wg sync.WaitGroup
	// Канал для передачи данных между воркерами и писателем
	ch := make(chan any)

	// Запуск горутины для писателя
	wg.Add(1)
	go func() {
		defer wg.Done()
		writer(ch, sigChan)
		signal.Stop(sigChan)
		close(ch)
	}()

	// Запуск воркеров
	for i := 0; i < workersCount; i++ {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			worker(ch, idx)
		}(i + 1)
	}

	// Ожидание завершения всех горутин
	wg.Wait()
}

// Функция writer генерирует случайные числа и отправляет их в канал
func writer(out chan<- any, sigChan <-chan os.Signal) {
	for {
		select {
		case sig := <-sigChan:
			// Обработка сигнала
			switch sig {
			// Получен сигнал прерывания (Ctrl+C)
			case os.Interrupt:
				return
			}
		default:
			out <- rand.Int()
		}
	}
}

// Функция worker обрабатывает данные из канала и выводит их на экран
func worker(in <-chan any, workerNum int) {
	for val := range in {
		fmt.Printf("Воркер №%d, Вывод: %v\n", workerNum, val)
	}
}
