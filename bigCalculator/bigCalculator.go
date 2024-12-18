package bigCalculator

import (
	"fmt"
	"math/big"
	"time"
)

func Run() {
	var num1Str, num2Str string
	fmt.Println("Digite o primeiro número (número grande):")
	fmt.Scan(&num1Str)
	fmt.Println("Digite o segundo número (número grande):")
	fmt.Scan(&num2Str)

	num1 := new(big.Int)
	num2 := new(big.Int)

	_, ok1 := num1.SetString(num1Str, 10)
	_, ok2 := num2.SetString(num2Str, 10)

	if !ok1 || !ok2 {
		fmt.Println("Erro: Um ou ambos os números são inválidos!")
		return
	}

	start := time.Now()

	result := new(big.Int)
	result.Add(num1, num2)

	duration := time.Since(start)

	fmt.Println("Resultado da soma de grandes números:", result)

	durationMicroseconds := float64(duration.Microseconds())
	durationMilliseconds := durationMicroseconds / 1000.0
	fmt.Printf("Tempo de execução: %.4f microssegundos (%.4f ms)\n", durationMicroseconds, durationMilliseconds)
}
