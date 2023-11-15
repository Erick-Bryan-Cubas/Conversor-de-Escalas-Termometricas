package main

import "fmt"

func main() {
	var escala float64
	fmt.Println("Escolha uma escala de temperatura: ")
	fmt.Println("1 - Celsius")
	fmt.Println("2 - Fahrenheit")
	fmt.Println("3 - Kelvin")
	fmt.Scan(&escala)

	var temperatura float64
	fmt.Println("Digite a temperatura: ")
	fmt.Scan(&temperatura)

	switch escala {
	case 1:
		fmt.Println("Celsius: ", temperatura)
		fmt.Println("Fahrenheit: ", temperatura*1.8+32)
		fmt.Println("Kelvin: ", temperatura+273.15)
	case 2:
		fmt.Println("Celsius: ", (temperatura-32)/1.8)
		fmt.Println("Fahrenheit: ", temperatura)
		fmt.Println("Kelvin: ", (temperatura-32)/1.8+273.15)
	case 3:
		fmt.Println("Celsius: ", temperatura-273.15)
		fmt.Println("Fahrenheit: ", (temperatura-273.15)*1.8+32)
		fmt.Println("Kelvin: ", temperatura)
	default:
		fmt.Println("Escala inválida")
	}

}
