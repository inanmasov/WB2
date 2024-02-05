package task6

import (
	"context"
	"fmt"
	"time"
)

// Функция Task6 реализует различные способы остановки горутин
func Task6() {
	// Использование канала для передачи сигнала
	done := make(chan bool)

	go func(done chan bool) {
		fmt.Println("1 вариант начало")
		time.Sleep(2 * time.Second)
		fmt.Println("1 вариант конец")
		done <- true
	}(done)

	// Ждем сигнала о завершении горутины
	<-done

	ctx, cancel := context.WithCancel(context.Background())

	// Использование контекста:
	go func(ctx context.Context) {
		fmt.Println("2 вариант начало")
		select {
		case <-time.After(2 * time.Second):
			fmt.Println("2 вариант конец")
		case <-ctx.Done():
			fmt.Println("Остановка 2 варианта")
		}
	}(ctx)

	// Остановка горутины с использованием функции cancel
	time.Sleep(1 * time.Second)
	cancel()

	done = make(chan bool)
	go func(done chan bool) {
		fmt.Println("3 вариант начало")
		time.Sleep(2 * time.Second)
		fmt.Println("3 вариант конец")
		done <- true
	}(done)

	// Использование select с каналом завершения
	select {
	case <-done:
		fmt.Println("Завершение 3 варианта")
	}

}
