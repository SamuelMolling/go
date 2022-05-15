package quadromensagens

import (
	"os"
	"sync"

	"github.com/jedib0t/go-pretty/v6/table"
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

var MsgStatusString = map[MsgStatus]string{ //map para tornar string
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
	Mensagem []*MsgMerc    // Numero máximo de mensagens do quadro
	MsgAtual int           //mensagem atual
	MuRW     *sync.RWMutex //Criar mutex de scrita e leitura
}

func (c *QuadroMsg) InicializaQmsg() { // Inicialização da estrutura de dados
	c.Mensagem = make([]*MsgMerc, 8)
	c.MsgAtual = 0
	c.MuRW = new(sync.RWMutex)
}

func (c *QuadroMsg) LivreQMsg() int { // retorna com o indice da mensagem livre
	for i := 0; i < 8; i++ {
		if c.Mensagem[i] == nil || c.Mensagem[i].Status == Livre { //valida se mensagem esta como nula ou livre
			return i //caso esteja retorna o indice
		}
	}

	return -1
}

func (c *QuadroMsg) SetQMsg(msg *MsgMerc) int { // Seta mensagem no quadro
	c.MuRW.Lock()
	index := c.LivreQMsg()
	if index != -1 {
		c.Mensagem[index] = msg
	}
	c.MuRW.Unlock()

	return index
}

func (c *QuadroMsg) PrintQMsg() { //Imprime quadro de mensagens

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)

	t.AppendHeader(table.Row{"Codigo Fornecedor", "Preco Venda", "Capacidade Fornecimento", "Codigo Comprador", "Demanda Solicitada", "Status"})

	for i := 0; i < 8; i++ {
		t.AppendRow([]interface{}{c.Mensagem[i].CodigoFornecedor, c.Mensagem[i].PrecoVenda, c.Mensagem[i].CapacidadeFornecimento, c.Mensagem[i].CodigoComprador, c.Mensagem[i].DemandaSolicitada, MsgStatusString[c.Mensagem[i].Status]})
	}

	t.Render()
}

func (m *MsgMerc) Clean() {
	m.CodigoFornecedor = -1
	m.PrecoVenda = 0
	m.CapacidadeFornecimento = 0
	m.CodigoComprador = -1
	m.DemandaSolicitada = 0
	m.Status = Livre
}
