package piCalculation

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func calculatePi(numPoints int, wg *sync.WaitGroup, results chan int) {
	defer wg.Done()
	var insideCircle int
	for i := 0; i < numPoints; i++ {
		x := rand.Float64()
		y := rand.Float64()
		if x*x+y*y <= 1 {
			insideCircle++
		}
	}
	results <- insideCircle
}

func displayProgress(current, total int) {
	progress := float64(current) / float64(total) * 100
	barLength := 50
	filledLength := int(float64(barLength) * (progress / 100))
	unfilledLength := barLength - filledLength

	progressBar := fmt.Sprintf("[%s%s] %.2f%%",
		string(repeat('#', filledLength)),
		string(repeat(' ', unfilledLength)),
		progress,
	)
	fmt.Printf("\r%s", progressBar)
}

func repeat(ch rune, count int) []rune {
	runes := make([]rune, count)
	for i := 0; i < count; i++ {
		runes[i] = ch
	}
	return runes
}

func Run() {
	rand.Seed(time.Now().UnixNano())

	numGoroutines := 100 // Número de goroutines a serem usadas
	numPoints := 1000000 // Número de pontos por goroutine
	totalPoints := numGoroutines * numPoints

	var wg sync.WaitGroup
	results := make(chan int, numGoroutines)

	// Contador para progresso
	var pointsProcessed int

	// Criar goroutines
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func() {
			calculatePi(numPoints, &wg, results)
			pointsProcessed += numPoints
			displayProgress(pointsProcessed, totalPoints) // Exibe o progresso
		}()
	}

	// Iniciar o tempo
	startTime := time.Now()

	// Esperar até todas as goroutines terminarem
	wg.Wait()
	close(results)

	// Somar os resultados de todas as goroutines
	var totalInsideCircle int
	for result := range results {
		totalInsideCircle += result
	}

	// Calcular o valor de pi
	pi := 4.0 * float64(totalInsideCircle) / float64(totalPoints)
	elapsedTime := time.Since(startTime)

	// Exibir resultados
	fmt.Printf("\nValor estimado de Pi: %.10f\n", pi) // Exibe Pi com 10 casas decimais
	fmt.Printf("Tempo de execução: %s\n", elapsedTime)
}
