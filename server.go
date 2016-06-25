package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"io"
	"log"
	"strings"
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

	escolhe(conn)

}
func escolhe(conn net.Conn) {
	leitor := bufio.NewReader(conn)
	//escreve na conexão
	conn.Write([]byte("EDITAR ou NOVO" + "\n"))
	//espera pelo enter do usuário
	mensagem, _ := leitor.ReadString('\n')
	mensagem = mensagem[0:len(mensagem)-2]
	if strings.EqualFold("EDITAR", mensagem){
		fmt.Printf("EDITAR" + "\n")
		lecod(conn)
	}else{
		fmt.Printf("NOVO" + "\n")
		salvaarquivo(conn)
	}
}

func lecod(conn net.Conn) {
	leitor := bufio.NewReader(conn)
	//escreve na conexão
	conn.Write([]byte("Digite o código" + "\n"))
	//espera pelo enter do usuário
	mensagem, _ := leitor.ReadString('\n')
	mensagem = mensagem[0:len(mensagem)-2]

    var path = mensagem+".txt"

    _, err := os.Stat(path)
    if err != nil {
        if os.IsNotExist(err) {
            log.Fatal("Arquivo não econtrado.")
        }
    }
    file, err := os.OpenFile(path, os.O_APPEND, 0666)
    if err != nil {
        log.Fatal(err)
    }
    // write some text to file
	_, err = file.WriteString("alterado")
	 if err != nil {
        log.Fatal(err)
    }

	// save changes
	err = file.Sync()
	if err != nil {
        log.Fatal(err)
    }
    file.Close()
}

func salvaarquivo(conn net.Conn) {
	leitor := bufio.NewReader(conn)
	//escreve na conexão
	conn.Write([]byte("Codigo" + "\n")) //IMPRIME A SOLICITACAO PARA DIGITAR O COD
	//apertar enter
	cod, _ := leitor.ReadString('\n') //COD RECEBE O QUE FOI INSERIDO
	
	
	//CRIACAO DO ARQUIVO
	filename := cod[0:len(cod)-2]+".txt" //DEFINE NOME DO ARQUIVO COMO SENDO COD, E REMOVE OS DOIS ULTIMOS CARACTERES (/N) 

     file, err := os.Create(filename) //DEFINE NOME DO ARQUIVO COMO COD+.txt

     if err != nil {
         fmt.Println(err) //IMPRIME ERRO NA TELA CASO ALGO DE ERRADO
     }

     fmt.Println("SALVAR NO ARQUIVO : " + filename) //MOSTRA QUAL O NOME DO ARQUIVO QUE SERA CRIADO
	 
	 n, err := io.WriteString(file, "Código: " + cod) //ESCREVE DADOS INSERIDOS NO PROD, PULA LINHA, ESCREVE DADOS INSERIDOS NO COD

     if err != nil {
         fmt.Println(n, err) //IMPRIME ERRO NA TELA CASO ALGO DE ERRADO
     }

     file.Close() //FECHA O ARQUIVO
 }
