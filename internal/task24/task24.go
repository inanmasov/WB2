package task24

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{x: x, y: y}
}

func dist(p1, p2 *Point) float64 {
	return math.Sqrt(math.Pow((p2.x-p1.x), 2) + math.Pow((p2.y-p1.y), 2))
}

// Реализация программы нахождения расстояния между двумя точками
func Task24() {
	var x1, y1 float64
	fmt.Print("Введите координату x для точки 1: ")
	fmt.Scan(&x1)
	fmt.Print("Введите координату y для точки 1: ")
	fmt.Scan(&y1)

	var x2, y2 float64
	fmt.Print("Введите координату x для точки 2: ")
	fmt.Scan(&x2)
	fmt.Print("Введите координату y для точки 2: ")
	fmt.Scan(&y2)

	p1 := NewPoint(x1, y1)
	p2 := NewPoint(x2, y2)
	fmt.Println(dist(p1, p2))
}
