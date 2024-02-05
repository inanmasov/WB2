package task8

import (
	"fmt"
	"strconv"
)

// Функция Task8 реализует изменение бита 64-битного числа
func Task8() {
	// Переменные для ввода данных.
	var num int64
	var pos uint
	var bit bool

	// Ввод данных.
	fmt.Print("Введите 64-битное число: ")
	fmt.Scan(&num)

	fmt.Print("Введите позицию бита: ")
	fmt.Scan(&pos)

	fmt.Print("Установить бит в 1 или 0 (true/false): ")
	fmt.Scan(&bit)

	// Преобразование в строку с двоичным представлением
	binaryStr := strconv.FormatInt(num, 2)
	fmt.Printf("%d в двоичной системе: %s\n", num, binaryStr)

	// Изменение бита
	num = setBit(num, pos, bit)

	// Преобразование в строку с двоичным представлением
	binaryStr = strconv.FormatInt(num, 2)
	fmt.Printf("%d в двоичной системе: %s\n", num, binaryStr)
}

func setBit(num int64, pos uint, bit bool) int64 {
	// Если исходный бит совпадает с заменяемым, ничего не меняется
	if 1&(num>>pos) == 1 {
		if bit {
			return num
		}
	} else {
		if !bit {
			return num
		}
	}
	// Ксорим бит с 1. Если был 0 станет 1. Если был 1 станет 0
	return num ^ (1 << pos)
}
