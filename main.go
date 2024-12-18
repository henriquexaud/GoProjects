package main

import (
	"fmt"
	"meu-projeto-go/bigCalculator"
	"meu-projeto-go/downloads"
	"meu-projeto-go/piCalculation"
	"meu-projeto-go/primes"
	"meu-projeto-go/webscraper"
	"os"
	"time"
)

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func showHeader() {
	clearScreen()
	fmt.Println("=========================================")
	fmt.Println("   🚀 Bem-vindo ao Demonstrador Go 🚀  ")
	fmt.Println("=========================================")
	fmt.Println()
}

func showMenu() {
	fmt.Println("Escolha uma opção abaixo:")
	fmt.Println("1 - Download concorrente")
	fmt.Println("2 - Cálculo de números primos")
	fmt.Println("3 - Calcular PI")
	fmt.Println("4 - Web Scraping")
	fmt.Println("5 - Cálculo de grande escala")
	fmt.Println("6 - Sair")
	fmt.Print("\nDigite sua escolha: ")
}

func main() {
	for {
		showHeader()
		showMenu()

		var choice int
		_, err := fmt.Scan(&choice)

		if err != nil {
			fmt.Println("\nErro: Entrada inválida! Por favor, digite um número.")
			time.Sleep(2 * time.Second)
			continue
		}

		switch choice {
		case 1:
			fmt.Println("\n🔄 Executando o exemplo de Download concorrente...")
			downloads.Run()
			fmt.Println("\n✅ Download concluído!")
		case 2:
			fmt.Println("\n🔢 Executando o cálculo de números primos...")
			primes.Run()
			fmt.Println("\n✅ Cálculo de números primos concluído!")
		case 3:
			fmt.Println("\n🌐 Executando o exemplo de Pi...")
			piCalculation.Run()
			fmt.Println("\n✅ Pi concluído!")
		case 4:
			fmt.Println("\n🌐 Executando o exemplo de Web Scraping...")
			webscraper.Run()
			fmt.Println("\n✅ Web Scraping concluído!")
		case 5:
			fmt.Println("\n🔢 Executando o cálculo de grande escala...")
			bigCalculator.Run()
			fmt.Println("\n✅ Cálculo de grande escala concluído!")
		case 6:
			fmt.Println("\n👋 Saindo do programa. Até logo!")
			os.Exit(0)
		default:
			fmt.Println("\n❌ Opção inválida! Tente novamente.")
		}

		fmt.Println("Pressione Enter para voltar ao menu principal...")
		fmt.Scanln()
		fmt.Scanln()
	}
}
