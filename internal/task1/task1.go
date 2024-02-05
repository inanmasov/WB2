package task1

import "fmt"

// Определение структуры Human
type Human struct {
	name    string
	age     int
	country string
}

// Определение структуры Action, встраивающей структуру Human
type Action struct {
	Human
}

// Метод получения имени
func (h Human) getName() string {
	return h.name
}

// Метод получения возраста
func (h Human) getAge() int {
	return h.age
}

// Метод получения страны
func (h Human) getCountry() string {
	return h.country
}

// Метод вывода всей информации
func (h Human) PrintInfo() {
	fmt.Printf("Name: %s, Age: %d, Country: %s\n", h.name, h.age, h.country)
}

// Функция Task1 для выполнения задачи.
func Task1() {
	human := Human{name: "Teylor", age: 22, country: "USA"}

	action := Action{Human: human}

	fmt.Println(action.getName())
	fmt.Println(action.getAge())
	fmt.Println(action.getCountry())
	action.PrintInfo()
}
