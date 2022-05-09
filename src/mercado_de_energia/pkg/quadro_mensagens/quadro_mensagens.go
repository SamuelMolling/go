package quadromensagens

import (
	"fmt"
	"sync"
)

// Variável Status da Mensagem, utilizada para controlar o quadro
type MsgStatus int

const (
	Livre    MsgStatus = iota // mensagem está livre para ser utilizada
	Oferta                    // oferta é postada pelo fornecedor
	Proposta                  // Comprador estabelece uma proposta de compra
	Aceite                    // Comprador aceita a proposta
	Recusa                    // Comprador nao aceita a proposta
)

var m = map[MsgStatus]string{ //map para tornar string
	Livre:    "Livre",
	Oferta:   "Oferta",
	Proposta: "Proposta",
	Aceite:   "Aceite",
	Recusa:   "Recusa",
}

// Estrutura de mensagens
type MsgMerc struct {
	CodigoFornecedor       int       // Código do fornecedor
	PrecoVenda             float64   // preco do kWh
	CapacidadeFornecimento float64   // capacidade de fornecimento, em kWh
	CodigoComprador        int       // codigo do comprador
	DemandaSolicitada      float64   // Demanda solicitada para contrato
	Status                 MsgStatus // status da negociação
}

type QuadroMsg struct { // Estrutura do quadro de mensagens
	Mensagem []MsgMerc     // Numero máximo de mensagens do quadro
	MsgAtual int           //mensagem atual
	MuRW     *sync.RWMutex //Criar mutex de scrita e leitura
}

func (c *QuadroMsg) InicializaQmsg() { // Inicialização da estrutura de dados
	c.Mensagem = make([]MsgMerc, 8)
	c.MuRW = new(sync.RWMutex)
}

func (c *QuadroMsg) LivreQMsg() int { // retorna com o indice da mensagem atual
	return c.MsgAtual
}

func (c *QuadroMsg) ProxQMsg() int { // Aponta para a proxima mensagem
	proxQMsg := c.LivreQMsg() + 1
	return proxQMsg
}

func (c *QuadroMsg) PrintQMsg() { //Imprime quadro de mensagens
	fmt.Printf("\n--------------------")
	for i := 0; i < 8; i++ {
		fmt.Printf("\n%s de energia", m[c.Mensagem[i].Status])
		fmt.Printf("\nQuadro %d", c.LivreQMsg())
		fmt.Printf("\nComprador %d", c.Mensagem[i].CodigoComprador)
		fmt.Printf("\nDemanda solicitada %.2f", c.Mensagem[i].DemandaSolicitada)
		fmt.Printf("\nComprador %d", c.Mensagem[i].Status)
		fmt.Printf("\n--------------------")
	}
}
