package task21

import "fmt"

// CelsiusThermometer - интерфейс для термометра, измеряющего температуру в градусах Цельсия
type CelsiusThermometer interface {
	GetTemperature() float64
}

// CelsiusThermometerImpl - реализация интерфейса CelsiusThermometer
type CelsiusThermometerImpl struct {
	temperatureCelsius float64
}

func (t *CelsiusThermometerImpl) GetTemperatureCelsius() float64 {
	return t.temperatureCelsius
}

// FahrenheitAdapter - адаптер для конвертации температуры из градусов Цельсия в градусы Фаренгейта
type FahrenheitAdapter struct {
	celsiusThermometer *CelsiusThermometerImpl
}

func (a *FahrenheitAdapter) GetTemperature() float64 {
	celsius := a.celsiusThermometer.GetTemperatureCelsius()
	fahrenheit := (celsius * 9.0 / 5.0) + 32.0
	return fahrenheit
}

// Реализация паттерна «адаптер»
func Task21() {
	// Создаем объект термометра, измеряющего температуру в градусах Цельсия
	celsiusThermometer := &CelsiusThermometerImpl{temperatureCelsius: 25.0}

	// Создаем адаптер для конвертации температуры из градусов Цельсия в градусы Фаренгейта
	var FahrenheitThermometer CelsiusThermometer = &FahrenheitAdapter{celsiusThermometer: celsiusThermometer}

	// Используем адаптер для получения температуры в градусах Фаренгейта
	temperatureFahrenheit := FahrenheitThermometer.GetTemperature()

	// Выводим результат
	fmt.Printf("Температура в градусах Фаренгейта: %.2f°F\n", temperatureFahrenheit)
}
