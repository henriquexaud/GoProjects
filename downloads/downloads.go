package downloads

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func downloadResource(id int, wg *sync.WaitGroup, results chan<- string) {
	defer wg.Done()

	duration := rand.Intn(3000) + 1000

	progressBar(id, 0, duration)

	for i := 0; i <= duration; i += 100 {
		time.Sleep(100 * time.Millisecond)
		progressBar(id, i, duration)
	}

	progressBar(id, duration, duration)

	results <- fmt.Sprintf("Baixado em %d ms", duration)
}

func progressBar(id, progress, total int) {
	percent := float64(progress) / float64(total) * 100
	barLength := 40

	bar := ""
	for i := 0; i < int(percent/2); i++ {
		bar += "#"
	}

	for i := len(bar); i < barLength; i++ {
		bar += " "
	}

	fmt.Printf("\rRecurso %d [%s] %.2f%%", id, bar, percent)
}

func Run() {
	fmt.Println("Executando o exemplo de download...")

	rand.Seed(time.Now().UnixNano())

	const totalResources = 5

	var wg sync.WaitGroup

	results := make(chan string, totalResources)

	fmt.Println("Iniciando download de recursos...")

	for i := 1; i <= totalResources; i++ {
		wg.Add(1)
		go downloadResource(i, &wg, results)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		fmt.Println("\n" + result)
	}

	fmt.Println("\nTodos os downloads foram concluídos!")
}
