package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitor = 5
const delay = 5

func main() {
	printIntro()

	for {
		showMenu()
		command := readCommand()

		switch command {
		case 1:
			startMonitoring()
		case 2:
			printLog()
		case 0:
			fmt.Println("Saindo...")
			os.Exit(0)
		default:
			fmt.Println("Não conheço este comando.")
		}
	}

}

func printIntro() {
	name := "Nico"
	version := 1.1
	fmt.Println("Olá sr.", name)
	fmt.Println("Este programa está na versão ", version)
}

func readCommand() int {
	var command int
	fmt.Scanf("%d", &command)
	return command
}

func showMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir logs")
	fmt.Println("0- Sair")
}

func startMonitoring() {
	fmt.Println("Monitorando...")
	// sites := []string{"https://www.alura.com.br", "https://google.com", "https://www.pechinchou.com.br", "https://www.caelum.com.br"}

	sites := readSitesFromFile()

	for i := 0; i < monitor; i++ {
		fmt.Println("Testando", i)
		for _, site := range sites {
			testSite(site)
		}
		fmt.Println("")
		time.Sleep(delay * time.Second)
	}
}

func testSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso.")
		registerLog(site, true)
	} else {
		fmt.Println("Site:", site, "está com problemas. Status", resp.StatusCode)
		registerLog(site, false)
	}
}

func readSitesFromFile() []string {
	var sites []string
	file, err := os.Open("sites.txt")
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	reader := bufio.NewReader(file)

	for {
		linha, err := reader.ReadString('\n')
		linha = strings.TrimSpace(linha)
		sites = append(sites, linha)
		if err == io.EOF {
			break
		}
	}
	file.Close()
	return sites
}

func registerLog(site string, online bool) {
	file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}
	file.WriteString(time.Now().Format("02/01/2006 15:04:05 - ") + site + " - online: " + strconv.FormatBool(online) + "\n")
	file.Close()
}

func printLog() {
	fmt.Println("Exibindo logs...")
	file, err := ioutil.ReadFile("log.txt")
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}
	fmt.Println(string(file))
}
