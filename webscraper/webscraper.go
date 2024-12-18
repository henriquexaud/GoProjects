package webscraper

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/gocolly/colly/v2"
)

func scrapeHeadlines(url string, siteName string, wg *sync.WaitGroup, results chan<- string) {
	defer wg.Done()

	c := colly.NewCollector()

	var title string

	c.OnHTML("h3, h2, .headline, .article__title", func(e *colly.HTMLElement) {
		title = e.Text
	})

	c.OnError(func(r *colly.Response, err error) {
		results <- fmt.Sprintf("Erro ao acessar '%s': %v", url, err)
	})

	err := c.Visit(url)
	if err != nil {
		results <- fmt.Sprintf("Erro ao acessar '%s': %v", url, err)
	}

	if title != "" {
		title = strings.TrimSpace(title)
		results <- fmt.Sprintf("%s - %s", siteName, title)
	}
}

func Run() {
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

	results := make(chan string, len(newsSites)*10)

	var wg sync.WaitGroup

	startTime := time.Now()

	for _, site := range newsSites {
		wg.Add(1)
		go scrapeHeadlines(site.url, site.siteName, &wg, results)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	fmt.Println("🔍 Buscando as principais notícias...")

	seen := make(map[string]bool)
	for result := range results {
		if len(seen) < 5 {
			if !seen[result] {
				fmt.Println("----------------------------------------------------")
				fmt.Println(result)
				fmt.Println("----------------------------------------------------")
				seen[result] = true
			}
		}
	}

	elapsedTime := time.Since(startTime)
	fmt.Printf("\nTempo total de execução: %s\n", elapsedTime)
}
