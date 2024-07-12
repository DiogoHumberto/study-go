package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

const qtMonitoramentos = 2
const delay = 5

func main() {

	exibeIntroducao()

	for {

		exibeMenu()

		comando := lerComando()

		//switch
		switch comando {
		case 1:
			initMonitoramento()
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

func initMonitoramento() {

	sites := leSitesArquivo()

	fmt.Println("Monitorando...")

	for i := 1; i <= qtMonitoramentos; i++ {
		fmt.Println("!!---> cliclo atual ", i, " de ", qtMonitoramentos)
		for _, site := range sites {
			testeSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}

}

func leSitesArquivo() []string {

	var sites []string
	arquivo, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	leitor := bufio.NewReader(arquivo)

	for {
		linha, err := leitor.ReadString('\n')

		linha = strings.TrimSpace(linha)

		sites = append(sites, linha)

		if err == io.EOF {
			break
		}
	}

	arquivo.Close()

	return sites
}

func testeSite(site string) {

	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("<<-->> Site:", site, ": foi carregado com sucesso!")
	} else {
		fmt.Println("!!!-->> Site:", site, " : esta com problemas. Status Code:", resp.StatusCode)
	}
}

// somente de estudo
func usoRetornos() {
	retorno1, retorno2 := metodoComDoisRetornos()
	fmt.Println(retorno1)
	fmt.Println(retorno2)

	//ignorando o segundo retorno
	retorno1Ignorado, _ := metodoComDoisRetornos()
	fmt.Println(retorno1Ignorado)
}

// retornos multiplos
func metodoComDoisRetornos() (string, string) {
	return "Retorno 1", "Retorno 2"
}

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

//alternativa abrir arquivo
//arq, err := os.Open("sites.txt") direto via sistema

// arq, err := ioutil.ReadFile("sites.txt") // via io/ioutil
// if err != nil {
// 	fmt.Println("Ocorreu um erro:", err)
// }
