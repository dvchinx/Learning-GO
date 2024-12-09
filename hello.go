package main

import "fmt"

var numero = 22
var numero2 = 5
var nombre string = "JESUS"

func main() {

	// Uso de metodos
	saludo("Hola, Go!")

	// Usar "+" para sumar dos variables
	fmt.Println(numero + numero2)

	// Usar "+" concatena las variables sin espacios
	fmt.Println("Mi nombre es: " + nombre)

	// Usar "," para separar los valores genera un espacio de 2
	fmt.Println("Mi nombre es: ", nombre, " y mi edad es: ", numero)

	if numero > numero2 {
		fmt.Println("El numero es mayor")
	}
}

func saludo(msg string) {
	fmt.Println(msg)
}
