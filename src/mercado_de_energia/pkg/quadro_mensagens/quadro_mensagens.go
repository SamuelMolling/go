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
	Oferta                    // oferta é postada pelo Comprador
	Proposta                  // Fornecedor estabelece uma proposta de compra
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
	Mensagem     []*MsgMerc // Numero máximo de mensagens do quadro
	MensagemLock []*sync.Mutex
	MsgAtual     int //mensagem atual
	Mu           *sync.Mutex
}

func (c *QuadroMsg) InicializaQmsg() { // Inicialização da estrutura de dados
	c.Mensagem = make([]*MsgMerc, 8)
	c.MensagemLock = make([]*sync.Mutex, 8)
	c.Mu = &sync.Mutex{}
	for i := 0; i < 8; i++ {
		c.Mensagem[i] = new(MsgMerc)
		c.Mensagem[i].Clean()
		c.MensagemLock[i] = new(sync.Mutex)
	}

	c.MsgAtual = 0
}

func (c *QuadroMsg) SetQMsg(msg *MsgMerc) int { // Seta mensagem no quadro
	c.Mu.Lock()
	defer c.Mu.Unlock()
	for i := 0; i < 8; i++ {
		if c.Mensagem[i].Status == Livre {
			c.Mensagem[i] = msg
			return i
		}
	}

	return -1
}

func (c *QuadroMsg) GetQMsg(codigoComprador int) int {
	for i := 0; i < 8; i++ {
		if c.Mensagem[i] != nil && c.Mensagem[i].CodigoComprador == codigoComprador {
			return i
		}
	}
	return -1
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

func (m *MsgMerc) CleanFornecedor() {
	m.CodigoFornecedor = -1
	m.PrecoVenda = 0
	m.CapacidadeFornecimento = 0
	m.Status = Oferta
}
