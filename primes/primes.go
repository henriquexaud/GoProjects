package primes

import (
	"fmt"
	"math"
	"sync"
	"time"
)

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func findPrimes(start, end int, wg *sync.WaitGroup, results chan<- int, progress *int, total int, mu *sync.Mutex) {
	defer wg.Done()

	for i := start; i <= end; i++ {
		if isPrime(i) {
			results <- i
		}

		mu.Lock()
		*progress++
		mu.Unlock()

		if *progress%100 == 0 {
			progressBar(*progress, total)
		}
	}
}

func progressBar(progress, total int) {
	percent := float64(progress) / float64(total) * 100
	barLength := 40

	bar := ""
	for i := 0; i < int(percent/2); i++ {
		bar += "#"
	}

	for i := len(bar); i < barLength; i++ {
		bar += " "
	}

	fmt.Printf("\rCalculando primos... [%s] %.2f%%", bar, percent)
}

func Run() {
	fmt.Println("Executando o exemplo de cálculo de primos...")
	startRange := 1
	endRange := 1000000
	workers := 4

	fmt.Printf("Calculando números primos entre %d e %d usando %d workers...\n", startRange, endRange, workers)

	startTime := time.Now()

	results := make(chan int, 100)
	var wg sync.WaitGroup

	var progress int

	var mu sync.Mutex

	rangeSize := (endRange - startRange + 1) / workers

	for i := 0; i < workers; i++ {
		wg.Add(1)
		start := startRange + i*rangeSize
		end := start + rangeSize - 1
		if i == workers-1 {
			end = endRange
		}

		go findPrimes(start, end, &wg, results, &progress, endRange, &mu)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	primes := []int{}
	for prime := range results {
		primes = append(primes, prime)
	}

	elapsedTime := time.Since(startTime)

	fmt.Printf("\nTotal de números primos encontrados: %d\n", len(primes))
	fmt.Printf("Tempo de execução: %s\n", elapsedTime)
}
