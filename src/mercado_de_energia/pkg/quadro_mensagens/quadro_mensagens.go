package quadromensagens

import (
	"fmt"
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

// Estrutura de mensagens
type MsgMerc struct {
	CodigoFornecedor       int       // Código do fornecedor
	PrecoVenda             float64   // preco do kWh
	CapacidadeFornecimento float64   // capacidade de fornecimento, em kWh
	CodigoComprador        int       // codigo do comprador
	DemandaSolicitada      float64   // Demanda solicitada para contrato
	Status                 MsgStatus // status da negociação
}

// Estrutura do quadro de mensagens
type QuadroMsg struct {
	Mensagem []MsgMerc // Numero máximo de mensagens do quadro
	MsgAtual int       //mensagem atual
}

// Inicialização da estrutura de dados
//void inicQMsg(QMsg *);
func (c *QuadroMsg) InicializaQmsg() {
	c.MsgAtual = 1
}

//int livreQMsg(QMsg *);
func (c *QuadroMsg) LivreQMsg() int { // retorna com o indice da mensagem
	return c.MsgAtual
}

func (c *QuadroMsg) proxQMsg() int { // Aponta para a proxima mensagem
	proxQMsg := c.LivreQMsg() + 1
	return proxQMsg
}

//FAZER RETURN STRING
func (c *QuadroMsg) PrintQMsg() { //Imprime quadro de mensagens
	fmt.Println(c.Mensagem[8])
}

func (c *MsgMerc) AtualizaDemandaQuadro(id int) {
	if c.Status == 3 {
		id = c.CodigoComprador //Validar necessidade
		//consumidor := consumidor.Efornecedor{Id: 1}
		//consumidor := "consumidor" + strconv.Itoa(id)

	}
}
