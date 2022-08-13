package main

import "fmt"

const number = 35

func celsiusToKalvin(c float64) float64 {
	temperature := c + 273.15
	return temperature
}

func celsiusToFahrenheit(f float64) float64 {
	temperature := (f * 9 / 5) + 32
	return temperature
}

func printTemperature(k, f float64) string {
	return fmt.Sprintf("Температура в Калвинах: %v°K\nТемпература в Фарингейтах: %v°F\n ", k, f)
}

func main() {

	kalvin := celsiusToKalvin(number)
	fahrenheit := celsiusToFahrenheit(number)

	result := printTemperature(kalvin, fahrenheit)
	fmt.Println(result)
}
