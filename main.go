package main

import "fmt"

type celcius struct {
	value float64
}

type fahrenheit struct {
	value float64
}

type kelvin struct {
	value float64
}

func (c celcius) toFahrenheit() float64 {
	return ((9.0 / 5.0) * c.value) + 32
}

func (c celcius) toKelvin() float64 {
	return c.value + 273.15
}

func (c celcius) toCelcius() float64 {
	return c.value
}

func (f fahrenheit) toCelcius() float64 {
	return (f.value - 32) * (5.0 / 9.0)
}

func (f fahrenheit) toKelvin() float64 {
	return (f.value + 459.67) * (5.0 / 9.0)
}

func (f fahrenheit) toFahrenheit() float64 {
	return f.value
}

func (k kelvin) toCelcius() float64 {
	return k.value - 273.15
}

func (k kelvin) toFahrenheit() float64 {
	return (k.value * (9.0 / 5.0)) - 459.67
}

func (k kelvin) toKelvin() float64 {
	return k.value
}

type calculateTemprature interface {
	toFahrenheit() float64
	toKelvin() float64
	toCelcius() float64
}

func main() {
	fmt.Println("Temprature Conversion")
	fmt.Println("1. Celcius")
	fmt.Println("2. Fahrenheit")
	fmt.Println("3. Kelvin")
	fmt.Println("Enter the temprature value:")

	var firstOption uint64
	fmt.Scanf("%d", &firstOption)

	for firstOption < 1 || firstOption > 3 {
		fmt.Println("Invalid value. Please enter valid value")
		fmt.Scanf("%d", &firstOption)
	}

	fmt.Println("Insert option to convert to:")
	fmt.Println("1. Celcius")
	fmt.Println("2. Fahrenheit")
	fmt.Println("3. Kelvin")
	fmt.Println("Enter the option value (1-3):")

	var secondOption uint64
	fmt.Scanf("%d", &secondOption)

	for secondOption < 1 || secondOption > 3 {
		fmt.Println("Invalid value. Please enter valid value")
		fmt.Scanf("%d", &secondOption)
	}

	var value float64
	fmt.Println("Enter the temprature value:")
	fmt.Scanf("%f", &value)

	var interfaceValue calculateTemprature
	switch firstOption {
	case 1:
		interfaceValue = celcius{value}
	case 2:
		interfaceValue = fahrenheit{value}
	case 3:
		interfaceValue = kelvin{value}
	}

	var result float64
	switch secondOption {
	case 1:
		result = interfaceValue.toCelcius()
	case 2:
		result = interfaceValue.toFahrenheit()
	case 3:
		result = interfaceValue.toKelvin()
	}

	fmt.Println("Result:", result)
}
