package task14

import (
	"fmt"
	"reflect"
)

// Функция Task14 реализует программу, которая в рантайме определяет тип переменной
func Task14() {
	// Примеры переменных разных типов
	var intValue int = 42
	var stringValue string = "Hello, World!"
	var float64Value float64 = 20.3
	var boolValue bool = true
	var chanValue chan int
	var pointer *int = nil

	// Определение типа переменных в рантайме
	printType(intValue)
	printType(stringValue)
	printType(float64Value)
	printType(boolValue)
	printType(chanValue)
	printType(pointer)
}

func printType(value interface{}) {
	// Используем пакет reflect для определения типа переменной
	switch reflect.TypeOf(value).Kind() {
	case reflect.Int:
		fmt.Println(value, "имеет тип: int")
	case reflect.String:
		fmt.Println(value, "имеет тип: string")
	case reflect.Float64:
		fmt.Println(value, "имеет тип: float64")
	case reflect.Bool:
		fmt.Println(value, "имеет тип: bool")
	case reflect.Chan:
		fmt.Println(value, "имеет тип: channel")
	default:
		fmt.Println("Неизвестный тип")
	}
}
