package task25

import (
	"fmt"
	"time"
)

func sleep1(sec int) {
	// Создаем канал для блокировки
	ch := make(chan bool)

	// Запускаем анонимную горутину для задержки
	go func() {
		time.Sleep(time.Duration(sec) * time.Second)
		ch <- true
	}()

	// Ожидаем завершения горутины
	<-ch

}

// Реализация собственной функции sleep
func Task25_1() {
	fmt.Println("Начало работы")

	// Используем собственную функцию sleep на 3 секунды
	sleep1(3)

	fmt.Println("Завершение работы после задержки")
}
