package bigCalculator

import (
	"fmt"
	"math/big"
	"time"
)

// Função que realiza um cálculo pesado (por exemplo, somar grandes números)
func Run() {
	// Solicitar ao usuário os dois números
	var num1Str, num2Str string
	fmt.Println("Digite o primeiro número (número grande):")
	fmt.Scan(&num1Str)
	fmt.Println("Digite o segundo número (número grande):")
	fmt.Scan(&num2Str)

	// Convertemos os números fornecidos para tipo big.Int
	num1 := new(big.Int)
	num2 := new(big.Int)

	// Tentamos converter as strings para números inteiros grandes
	_, ok1 := num1.SetString(num1Str, 10)
	_, ok2 := num2.SetString(num2Str, 10)

	if !ok1 || !ok2 {
		fmt.Println("Erro: Um ou ambos os números são inválidos!")
		return
	}

	// Começo do cálculo
	start := time.Now()

	// Realizando a soma
	result := new(big.Int)
	result.Add(num1, num2)

	// Tempo de execução
	duration := time.Since(start)

	// Exibindo os resultados
	fmt.Println("Resultado da soma de grandes números:", result)

	// Exibindo o tempo de execução em microssegundos
	durationMicroseconds := float64(duration.Microseconds())
	durationMilliseconds := durationMicroseconds / 1000.0
	fmt.Printf("Tempo de execução: %.4f microssegundos (%.4f ms)\n", durationMicroseconds, durationMilliseconds)
}
