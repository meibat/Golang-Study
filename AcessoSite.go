package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
	"io"
	"bufio"
)

const delay = 3

func main() {
	for {
		opcao := introducao()

		switch opcao {
		case 1:
			var numero int
			var site int
			sites := listarSites("sites.txt")

			fmt.Println("\nQual site deseja testar?")
			fmt.Scan(&site)
			fmt.Println("Quantas vezes deseja testar?")
			fmt.Scan(&numero)

			monitoramento(numero, site, sites)
		case 2:
			println("logs")
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

func listarSites(nomeArquivo string) []string{
	var sites []string
	arquivo, err := os.Open(nomeArquivo) 
	//arquivo, err := ioutil.ReadFile(nomeArquivo) retorna em byte

	if err != nil{
		fmt.Println("Ocorreu um erro: ", err)
		var nulo []string
		return nulo
	}

	leitor := bufio.NewReader(arquivo)
	i := 1
	for {
		linha, err := leitor.ReadString('\n')
		fmt.Print(i,"-", linha)
		sites = append(sites, linha)

		if err == io.EOF{
			break
		}
		i += 1
	}

	return sites
}

func monitoramento(numero int, site int, sites []string) {
	if site != 0 {
		for vezes := 0; vezes < numero; vezes++ {
			acessandoSite(sites[site])
			fmt.Println("-----------------------------")
			time.Sleep(delay * time.Second)
		}
	} else{
		for vezes := 0; vezes < numero; vezes++ {
			for i := 0; i < len(sites); i++ {
				acessandoSite(sites[i])
				fmt.Println("-----------------------------")
			}
			time.Sleep(delay * time.Second)
		}
	}
}

func acessandoSite(site string) {
	fmt.Println("Testando site: ", site)

	resp, err := http.Get(site)
	fmt.Println("Site Status:", resp.StatusCode)

	if err != nil || resp.StatusCode != 200 {
		fmt.Println("Site com problema! Erro:", err)
	} else {
		fmt.Println("O site foi acessado com sucesso!")
	} 
}
