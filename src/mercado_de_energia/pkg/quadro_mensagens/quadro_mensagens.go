// Numero máximo de mensagens do quadro
#define MAXMSG	8
	
// Variável Status da Mensagem, utilizada para controlar o quadro
const (
	Livre 	int = 0
	Oferta		= 1
	Proposta	= 2	
	Aceite		= 3
	Recusa		= 4
			
)

type MsgStatus enum {Livre=0,	// mensagem está livre para ser utilizada
		Oferta=1,	// oferta é postada pelo fornecedor
   	    Proposta=2, // Comprador estabelece uma proposta de compra
	    Aceite=3,	// Comprador aceita a proposta
	    Recusa=4	// Comprador nao aceita a proposta
} 

// Estrutura de mensagens
typedef struct {    
int          Fornecedor; // Código do fornecedor
double       pVenda;     // preco do kWh
double       CapForn;    // capacidade de fornecimento, em kWh
int          Comprador;   // codigo do comprador
double 	     DemSolic;   // Demanda solicitada para contrato
MsgStatus    Status;	   // status da negociação
}MsgMerc; 

	      
// Estrutura do quadro de mensagens
typedef struct{
MsgMerc	Msg[MAXMSG];
int	MsgAtual;
}QMsg;	

// Prototipos de funcoes
// Inicialização da estrutura de dados
void inicQMsg(QMsg *);

// Aponta para proxima mensagem livre
// retorna com o indice da mensagem
int livreQMsg(QMsg *);

// Aponta para a proxima mensagem
// retorna com o indice da mensagem
int proxQMsg(QMsg *);

//Imprime quadro de mensagens
void printQMsg(QMsg *);
