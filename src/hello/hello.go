package main

import (
	"fmt"
	"reflect"
)

func main() {
	nome := "Samuel"
	idade := 24
	versao := 1.1
	fmt.Println("Hello world", nome, "sua idade é,", idade)
	fmt.Println("Versão: ", versao)

	fmt.Println("Tipo da variável nome é:", reflect.TypeOf(idade))
}
