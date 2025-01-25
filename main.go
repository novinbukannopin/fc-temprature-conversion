package main

import "fmt"

func main() {
	fmt.Println("Masukkan kalkulator suhu yang ingin digunakan : ")
	fmt.Println("1. Celcius ke Fahrenheit")
	fmt.Println("2. Celcius ke Kelvin")
	fmt.Println("3. Fahrenheit ke Celcius")
	fmt.Println("4. Fahrenheit ke Kelvin")
	fmt.Println("5. Kelvin ke Celcius")
	fmt.Println("6. Kelvin ke Fahrenheit")
	fmt.Println("Pilih angka 1-6 : ")

	var option int
	fmt.Scanf("%d", &option)
	for option < 1 || option > 6 {
		fmt.Println("Pilihan tidak valid, silahkan pilih angka 1-6 : ")
		fmt.Scanf("%d", &option)
	}

	var value float64
	fmt.Println("Masukkan nilai suhu : ")
	fmt.Scanf("%f", &value)

	var result float64
	if option == 1 {
		result = CelciusToFahrenheit(value)
	} else if option == 2 {
		result = CelciusToKelvin(value)
	} else if option == 3 {
		result = FahrenheitToCelcius(value)
	} else if option == 4 {
		result = FahrenheitToKelvin(value)
	} else if option == 5 {
		result = KelvinToCelcius(value)
	} else {
		result = KelvinToFahrenheit(value)
	}

	fmt.Printf("Hasil konversi suhu adalah : %.2f\n", result)
}

func CelciusToFahrenheit(c float64) float64 {
	f := (9.0 / 5.0 * c) + 32
	return f
}

func CelciusToKelvin(c float64) float64 {
	k := c + 273.15
	return k
}

func FahrenheitToCelcius(f float64) float64 {
	c := (f - 32.0) * 5 / 9
	return c
}

func FahrenheitToKelvin(f float64) float64 {
	k := (f + 459.67) * (5 / 9)
	return k
}

func KelvinToCelcius(k float64) float64 {
	c := k - 273.15
	return c
}

func KelvinToFahrenheit(k float64) float64 {
	f := (k * (5 / 9)) - 459.67
	return f
}
