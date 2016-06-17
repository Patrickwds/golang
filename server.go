package main

import (
	"bufio"
	"fmt"
	"net"
)

// "stri

//link de instalação do golang
//https://golang.org/dl/
// depois de instalar, abrir o cmd e digitar
//C:\Go\bin\go.exe \caminho\para\o\arquivo\server.go

func main() {

	fmt.Println("Launching server...")
	ln, _ := net.Listen("tcp", ":8081") //define porta do server
	conn, _ := ln.Accept()

	lenome(conn)
	lecpf(conn)

}

func lenome(conn net.Conn) {
	leitor := bufio.NewReader(conn)
	//escreve na conexão
	conn.Write([]byte("Digite seu nome" + "\n"))
	//espera pelo enter do usuário
	mensagem, _ := leitor.ReadString('\n')
	//imprime o que foi digitado pelo usuario
	fmt.Printf("O Nome do cliente é:" + mensagem)
}

func lecpf(conn net.Conn) {
	leitor := bufio.NewReader(conn)
	//escreve na conexão
	conn.Write([]byte("Digite seu cpf" + "\n"))
	//espera pelo enter do usuário
	mensagem, _ := leitor.ReadString('\n')
	//imprime o que foi digitado pelo usuario
	fmt.Printf("O cpf do cliente é :" + mensagem)
}
