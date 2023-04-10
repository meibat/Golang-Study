package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const delay = 3

func main() {
	for {
		opcao := introducao()

		switch opcao {
		case 1:
			var numero int
			fmt.Println("Quantas vezes deseja testar?")
			fmt.Scan(&numero)
			monitoramento(numero)
		case 2:
			print("logs")
		case 0:
			fmt.Println("Encerrando o programa...")
			os.Exit(0)
		default:
			fmt.Println("Opção desconhecida. Encerrando...")
			os.Exit(-1)
		}
	}
}

func introducao() int {
	fmt.Println("Escolha uma opção:")
	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair")

	var opcao int
	fmt.Scan(&opcao)

	return opcao
}

func monitoramento(numero int) {
	sites := []string{"https://www.alura.com.br", "https://random-status-code.herokuapp.com/"} //slice

	for vezes := 0; vezes < numero; vezes++ {
		for i := 0; i < len(sites); i++ {
			acessandoSite(sites[i])
			fmt.Println("-----------------------------")
		}
		time.Sleep(delay * time.Second)
	}
}

func acessandoSite(site string) {
	fmt.Println("Testando site: ", site)
	resp, _ := http.Get(site)
	fmt.Println("Site Status:", resp.StatusCode)

	if resp.StatusCode == 200 {
		fmt.Println("O site foi acessado com sucesso!")
	} else {
		fmt.Println("Site com problema!")
	}
}
