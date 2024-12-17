package webscraper

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/gocolly/colly/v2"
)

// scrapeHeadlines busca as principais notícias de um canal de notícias
func scrapeHeadlines(url string, siteName string, wg *sync.WaitGroup, results chan<- string) {
	defer wg.Done()

	// Cria um novo coletor
	c := colly.NewCollector()

	// Variável para título
	var title string

	// Define o que fazer quando o título das notícias for encontrado
	c.OnHTML("h3, h2, .headline, .article__title", func(e *colly.HTMLElement) {
		title = e.Text
	})

	// Define o que fazer em caso de erro
	c.OnError(func(r *colly.Response, err error) {
		results <- fmt.Sprintf("Erro ao acessar '%s': %v", url, err)
	})

	// Visita o URL
	err := c.Visit(url)
	if err != nil {
		results <- fmt.Sprintf("Erro ao acessar '%s': %v", url, err)
	}

	// Se o título estiver presente, envia o dado
	if title != "" {
		// Remove espaços extras
		title = strings.TrimSpace(title)
		results <- fmt.Sprintf("%s - %s", siteName, title)
	}
}

// Run executa o exemplo de web scraping buscando notícias de sites populares brasileiros
func Run() {
	// Lista de URLs de sites de notícias e seus nomes
	newsSites := []struct {
		url      string
		siteName string
	}{
		{"https://g1.globo.com/", "G1 - Globo"},                 // G1 Globo
		{"https://www.folha.uol.com.br/", "Folha de S. Paulo"},  // Folha de S. Paulo
		{"https://www.uol.com.br/", "UOL"},                      // UOL
		{"https://www.estadao.com.br/", "O Estado de S. Paulo"}, // O Estado de S. Paulo
		{"https://www.valor.com.br/", "Valor Econômico"},        // Valor Econômico
		{"https://www.r7.com/", "R7"},                           // R7
	}

	// Canal para coletar os resultados das notícias
	results := make(chan string, len(newsSites)*10)

	// WaitGroup para aguardar todas as goroutines
	var wg sync.WaitGroup

	// Marca o tempo de início
	startTime := time.Now()

	// Inicia a coleta de dados de forma concorrente
	for _, site := range newsSites {
		wg.Add(1) // Incrementa o contador do WaitGroup
		go scrapeHeadlines(site.url, site.siteName, &wg, results)
	}

	// Goroutine para fechar o canal após todas as goroutines terminarem
	go func() {
		wg.Wait()
		close(results)
	}()

	// Exibe os resultados enquanto recebe dados do canal
	fmt.Println("🔍 Buscando as principais notícias...")

	// Exibe até 5 notícias de cada site
	seen := make(map[string]bool)
	for result := range results {
		if len(seen) < 5 {
			if !seen[result] {
				// Exibe as notícias de forma organizada
				fmt.Println("----------------------------------------------------")
				fmt.Println(result)
				fmt.Println("----------------------------------------------------")
				seen[result] = true
			}
		}
	}

	// Marca o tempo de fim e exibe o tempo total
	elapsedTime := time.Since(startTime)
	fmt.Printf("\nTempo total de execução: %s\n", elapsedTime)
}
