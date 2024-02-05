package main

import (
	"fmt"

	"example.com/2/internal/task1"
	"example.com/2/internal/task10"
	"example.com/2/internal/task11"
	"example.com/2/internal/task12"
	"example.com/2/internal/task13"
	"example.com/2/internal/task14"
	"example.com/2/internal/task15"
	"example.com/2/internal/task16"
	"example.com/2/internal/task17"
	"example.com/2/internal/task18"
	"example.com/2/internal/task19"
	"example.com/2/internal/task2"
	"example.com/2/internal/task20"
	"example.com/2/internal/task21"
	"example.com/2/internal/task22"
	"example.com/2/internal/task23"
	"example.com/2/internal/task24"
	"example.com/2/internal/task25"
	"example.com/2/internal/task26"
	"example.com/2/internal/task3"
	"example.com/2/internal/task4"
	"example.com/2/internal/task5"
	"example.com/2/internal/task6"
	"example.com/2/internal/task7"
	"example.com/2/internal/task8"
	"example.com/2/internal/task9"
)

func main() {
	for {
		var taskNumber int
		fmt.Print("Введите номер задания (1-26): ")
		fmt.Scan(&taskNumber)
		switch taskNumber {
		case 1:
			task1.Task1()
		case 2:
			task2.Task2_1()
			task2.Task2_2()
		case 3:
			task3.Task3()
		case 4:
			task4.Task4()
		case 5:
			task5.Task5()
		case 6:
			task6.Task6()
		case 7:
			task7.Task7()
		case 8:
			task8.Task8()
		case 9:
			task9.Task9()
		case 10:
			task10.Task10()
		case 11:
			task11.Task11()
		case 12:
			task12.Task12()
		case 13:
			task13.Task13()
		case 14:
			task14.Task14()
		case 15:
			task15.Task15()
		case 16:
			task16.Task16()
		case 17:
			task17.Task17()
		case 18:
			task18.Task18_1()
			task18.Task18_2()
		case 19:
			task19.Task19()
		case 20:
			task20.Task20()
		case 21:
			task21.Task21()
		case 22:
			task22.Task22()
		case 23:
			task23.Task23()
		case 24:
			task24.Task24()
		case 25:
			task25.Task25_1()
			task25.Task25_2()
		case 26:
			task26.Task26_1()
			task26.Task26_2()
		}

	}
}
