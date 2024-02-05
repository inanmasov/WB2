package task25

import (
	"fmt"
	"time"
)

func sleep2(sec int) {
	// Создаем канал для блокировки
	ch := make(chan bool)

	go func() {
		// Запускаем таймер и ждем указанное количество миллисекунд
		timer := time.NewTimer(time.Duration(sec) * time.Second)
		defer timer.Stop()

		// Ожидаем события от таймера или закрытия канала
		select {
		case <-timer.C:
		case <-ch:
		}
		close(ch)
	}()

	// Ожидаем завершения горутины
	<-ch
}

// Реализация собственной функции sleep
func Task25_2() {
	fmt.Println("Начало работы")

	// Используем собственную функцию sleep на 3 секунды
	sleep2(3)

	fmt.Println("Завершение работы после задержки")
}
