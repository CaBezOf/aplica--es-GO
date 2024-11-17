package main

import (
	"fmt"
)

func main() {
	// Variables for scores and weights
	var linguagens, humanas, natureza, matematica, redacao float64
	var pesoLinguagens, pesoHumanas, pesoNatureza, pesoMatematica, pesoRedacao float64

	// Defining the weights
	pesoLinguagens = 2
	pesoHumanas = 2
	pesoNatureza = 3
	pesoMatematica = 3
	pesoRedacao = 1

	// Input scores
	fmt.Println("Enter the score for Linguagens, Códigos e suas Tecnologias:")
	fmt.Scanln(&linguagens)

	fmt.Println("Enter the score for Ciências Humanas e suas Tecnologias:")
	fmt.Scanln(&humanas)

	fmt.Println("Enter the score for Ciências da Natureza e suas Tecnologias:")
	fmt.Scanln(&natureza)

	fmt.Println("Enter the score for Matemática e suas Tecnologias:")
	fmt.Scanln(&matematica)

	fmt.Println("Enter the score for Redação:")
	fmt.Scanln(&redacao)

	// Calculating the weighted average
	somaPesos := pesoLinguagens + pesoHumanas + pesoNatureza + pesoMatematica + pesoRedacao
	notaFinal := (linguagens*pesoLinguagens + humanas*pesoHumanas + natureza*pesoNatureza + matematica*pesoMatematica + redacao*pesoRedacao) / somaPesos

	// Displaying the result
	fmt.Printf("A sua média no ENEM é: %.2f\n", notaFinal)
}
