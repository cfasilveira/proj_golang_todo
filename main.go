package main

import (
	"fmt"
	"strings"
)

func main() {

	opcoes := []string{"Watch Go Crash Course", "Watch Nana's Golang Full Course", "Reward myself with a cheesecake"}

	fmt.Println(strings.Repeat("=", 20) + strings.ToUpper("| Welcome to our TodoList App! |") + strings.Repeat("=", 20))

	// Chamando a funcao printTasks para printar as opcoes de forma mais organizada
	printTasks(opcoes)

	fmt.Println("\nTutoriais:")
	fmt.Println(opcoes[0])
	fmt.Println(opcoes[1])

	fmt.Println("\nRewards:")
	fmt.Println(opcoes[2])

	fmt.Println("\nMy Projects:")
	fmt.Println(opcoes[1])

}

// Funcao para printar opcoes de forma mais organizada
func printTasks(task []string) {
	fmt.Println("\nList of my Todos:")
	for i, opcao := range task {
		fmt.Printf("%d. %s\n", i+1, strings.ToUpper(opcao))
	}
}
