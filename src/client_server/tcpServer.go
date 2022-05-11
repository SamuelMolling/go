package main

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

func main() {

	l, err := net.Listen("tcp", ":1234") //Listen na porta 1234, usando protocolo tcp
	if err != nil {                      //valida se não deu erro
		fmt.Println(err) //printa o erro caso dê erro
		return
	}
	defer l.Close() //fecha a validação

	c, err := l.Accept() //Aceita conexões
	if err != nil {      //valida se não deu erro
		fmt.Println(err) //printa o erro caso dê erro
		return
	}

	for { //for infinito
		netData, err := bufio.NewReader(c).ReadString('\n') //recebe conexão
		if err != nil {                                     //valida se não deu erro
			fmt.Println(err) //printa o erro caso dê erro
			return
		}

		fmt.Print("Message Receive -> ", string(netData)) //escreve a mensagem recebida pelo cliente
		t := time.Now()                                   //monta o timestamp da mensagem recebida
		myTime := t.Format(time.RFC3339) + "\n"
		c.Write([]byte(myTime)) //manda mensagem para o cliente com o timestamp
	}
}
