package downloads

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Simula o download de um recurso, aguardando um tempo aleatório
func downloadResource(id int, wg *sync.WaitGroup, results chan<- string) {
	defer wg.Done() // Marca a goroutine como concluída quando terminar

	// Simula um tempo de "download" entre 1 a 3 segundos
	duration := rand.Intn(3000) + 1000

	// Barra de progresso
	progressBar(id, 0, duration) // Inicia a barra de progresso

	for i := 0; i <= duration; i += 100 { // Atualiza a barra a cada 100ms
		time.Sleep(100 * time.Millisecond) // Espera um pouco
		progressBar(id, i, duration)       // Atualiza a barra
	}

	// Finaliza a barra de progresso
	progressBar(id, duration, duration)

	// Envia o resultado para o canal
	results <- fmt.Sprintf("Baixado em %d ms", duration)
}

// Exibe a barra de progresso no terminal
func progressBar(id, progress, total int) {
	// Calcula o percentual de progresso
	percent := float64(progress) / float64(total) * 100
	barLength := 40 // Tamanho da barra de progresso

	// Cria a barra
	bar := ""
	for i := 0; i < int(percent/2); i++ {
		bar += "#" // Cada "#" representa 2% de progresso
	}

	// Preenche o restante da barra com espaços
	for i := len(bar); i < barLength; i++ {
		bar += " "
	}

	// Limpa a linha atual e exibe o progresso
	fmt.Printf("\rRecurso %d [%s] %.2f%%", id, bar, percent)
}

// Função principal para executar os downloads
func Run() {
	fmt.Println("Executando o exemplo de download...")

	// Inicializa o gerador de números aleatórios
	rand.Seed(time.Now().UnixNano())

	const totalResources = 5

	// WaitGroup para sincronizar as goroutines
	var wg sync.WaitGroup

	// Canal para coletar os resultados dos downloads
	results := make(chan string, totalResources)

	fmt.Println("Iniciando download de recursos...")

	// Lança goroutines para download concorrente
	for i := 1; i <= totalResources; i++ {
		wg.Add(1) // Incrementa o contador do WaitGroup
		go downloadResource(i, &wg, results)
	}

	// Goroutine para fechar o canal quando os downloads terminarem
	go func() {
		wg.Wait() // Espera todas as goroutines terminarem
		close(results)
	}()

	// Recebe e exibe os resultados dos downloads
	for result := range results {
		fmt.Println("\n" + result)
	}

	fmt.Println("\nTodos os downloads foram concluídos!")
}
