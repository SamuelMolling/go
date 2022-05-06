package quadromensagens

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

// Aponta para proxima mensagem livre
// retorna com o indice da mensagem
//int livreQMsg(QMsg *);

// Aponta para a proxima mensagem
// retorna com o indice da mensagem
//int proxQMsg(QMsg *);

//Imprime quadro de mensagens
//func printQMsg(QMsg *) return string;
