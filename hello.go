package main

import (
	"fmt"
	"os"
)

func main() {

	exibeIntroducao()

	exibeMenu()

	comando := lerComando()

	//condicionais IF
	// if comando == 1 {
	// 	fmt.Println("Iniciando monitoramento...")
	// } else if comando == 2 {
	// 	fmt.Println("Exibindo logs...")
	// } else if comando == 0 {
	// 	fmt.Println("Saindo do programa...")
	// } else {
	// 	fmt.Println("Comando inválido")
	// }

	//switch
	switch comando {
	case 1:
		fmt.Println("Iniciando monitoramento...")
	case 2:
		fmt.Println("Exibindo logs...")
	case 0:
		fmt.Println("Saindo do programa...")
		os.Exit(0)
	default:
		fmt.Println("Comando inválido")
		os.Exit(-1)
	}
}

func exibeMenu() {
	fmt.Println("1- Iniciar monitoramento")
	fmt.Println("2- Exibir logs")
	fmt.Println("0- Sair do programa")
}

func exibeIntroducao() {
	//variaveis
	nome := "Diogo" // formato de declarar uma variavel
	idade := 25.0
	var versao float32 = 1.1 // pode ser atribuido automaticamente pelo Go

	fmt.Println("Olá amigo ", nome, "sua idade é ", idade, "sua versão é ", versao)
}

func lerComando() int {
	var comando int
	fmt.Scan(&comando)
	fmt.Println("--> Comando escolhido: ", comando)

	return comando
}
