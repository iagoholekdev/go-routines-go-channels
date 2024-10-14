package main

import (
	"fmt"
	"time"
)

// Função que envia números para o canal
func enviarNumeros(canal chan<- int) {
	for i := 1; i <= 3; i++ {
		fmt.Printf("Enviando: %d\n", i)
		canal <- i                  // Envia o número e espera até que o for no main receba
		time.Sleep(2 * time.Second) // Simula tempo de processamento
	}
	close(canal) // Fecha o canal
	fmt.Println("Canal fechado.")
}

func main() {
	// Cria um canal para enviar inteiros
	canal := make(chan int)

	// Inicia uma goroutine para enviar números
	go enviarNumeros(canal)

	// Lê e imprime os números recebidos do canal
	for numero := range canal {
		fmt.Printf("Recebido: %d\n", numero)
	}

	fmt.Println("Todos os números foram recebidos.")
}
