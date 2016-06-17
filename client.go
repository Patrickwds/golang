package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() { // connect to this socket

	conn, err := net.Dial("tcp", "localhost:8081")
	if err != nil {
		fmt.Print(err)
	}

	//abre leitor de conexao
	connectionreader := bufio.NewReader(conn)
	//espera por um enter
	mensagem, _ := connectionreader.ReadString('\n')
	//imprime o que o servidor mandou
	fmt.Print(mensagem)
	//abre o leitor de entrada de terminal
	osreader := bufio.NewReader(os.Stdin)
	//espera a entrada de texto e  armazena dentro de text
	text, _ := osreader.ReadString('\n')
	//envia pro servidor junto com enter
	fmt.Fprintf(conn, text+"\n")
	//espera resposta do server
	mensagem, _ = connectionreader.ReadString('\n')
	//printa resposta do server
	fmt.Print(mensagem)
	//le input do terminal
	text, _ = osreader.ReadString('\n')
	//printa na conexão o input do terminal
	fmt.Fprintf(conn, text)
}
