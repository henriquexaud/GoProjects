package primes

import (
	"fmt"
	"math"
	"sync"
	"time"
)

// isPrime verifica se um número é primo
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

// findPrimes encontra números primos em um intervalo específico
func findPrimes(start, end int, wg *sync.WaitGroup, results chan<- int, progress *int, total int, mu *sync.Mutex) {
	defer wg.Done() // Marca a goroutine como concluída quando terminar

	for i := start; i <= end; i++ {
		if isPrime(i) {
			results <- i // Envia o número primo para o canal
		}

		// Atualiza o progresso global
		mu.Lock()
		*progress++
		mu.Unlock()

		// Atualiza a barra de progresso a cada 100 números verificados
		if *progress%100 == 0 {
			progressBar(*progress, total)
		}
	}
}

// Exibe a barra de progresso no terminal
func progressBar(progress, total int) {
	percent := float64(progress) / float64(total) * 100
	barLength := 40 // Tamanho da barra de progresso

	bar := ""
	for i := 0; i < int(percent/2); i++ {
		bar += "#" // Cada "#" representa 2% de progresso
	}

	// Preenche o restante da barra com espaços
	for i := len(bar); i < barLength; i++ {
		bar += " "
	}

	// Limpa a linha atual e exibe o progresso
	fmt.Printf("\rCalculando primos... [%s] %.2f%%", bar, percent)
}

func Run() {
	fmt.Println("Executando o exemplo de cálculo de primos...")
	startRange := 1
	endRange := 1000000 // Define o intervalo de números
	workers := 4        // Número de goroutines (paralelismo)

	fmt.Printf("Calculando números primos entre %d e %d usando %d workers...\n", startRange, endRange, workers)

	// Marca o tempo de início
	startTime := time.Now()

	// Canal para coletar os números primos encontrados
	results := make(chan int, 100)
	var wg sync.WaitGroup

	// Contador global de progresso
	var progress int

	// Mutex para sincronizar o acesso ao contador de progresso
	var mu sync.Mutex

	// Divide o intervalo entre as goroutines (workers)
	rangeSize := (endRange - startRange + 1) / workers

	for i := 0; i < workers; i++ {
		wg.Add(1)
		start := startRange + i*rangeSize
		end := start + rangeSize - 1
		if i == workers-1 {
			end = endRange // Garante que o último worker vá até o final
		}

		go findPrimes(start, end, &wg, results, &progress, endRange, &mu)
	}

	// Goroutine para fechar o canal após todas as goroutines terminarem
	go func() {
		wg.Wait()
		close(results)
	}()

	// Coleta os resultados dos números primos
	primes := []int{}
	for prime := range results {
		primes = append(primes, prime)
	}

	// Marca o tempo de fim
	elapsedTime := time.Since(startTime)

	// Exibe os resultados
	fmt.Printf("\nTotal de números primos encontrados: %d\n", len(primes))
	fmt.Printf("Tempo de execução: %s\n", elapsedTime)
}
