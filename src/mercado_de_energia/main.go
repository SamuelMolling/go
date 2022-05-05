package main

import (
	"fmt"

)

typedef struct {
	double pD;    			// preco desejavel
	double miNpD;  			// menor preco admissivel
	double capacidade_carga;	// capacidade de carga
	double energia_gerada; 		// energia gerada
	double energia_fornecida; 	// energia fornecida
	double demanda_interna;		// demanda interna
	int faz_oferta;			//variavel que indica se fornecedor fez uma oferta (se ==1)
	}struct_Efornecedor; 
	
	//==============================================================================================================================
	//	Protópipos 
	//==============================================================================================================================
	
	// Inicialização da estrutura de dados
	void inicia_struct_Efornecedor(struct_Efornecedor *, int);
	
	// Atualizacao do pD
	void atualiza_pd(struct_Efornecedor *, double, double, double);
	
	//Gerador de num aleatorios
	double num_aleat (double, double);


	typedef struct {
		double tD;     			// tarifa desejavel
		double maXtD;  			// maximo preco admissivel
		double tC;			// prazo do contrato do cliente
		double demanda; 		// demanda do cliente
		}struct_EConsumidor; 
		
		//==============================================================================================================================
		//	Protópipos 
		//==============================================================================================================================
		
		// Inicialização da estrutura de dados
		void inicia_struct_EConsumidor(struct_EConsumidor *, int);
		
		//Atualiza tA
		void atualiza_pa(struct_EConsumidor *, double, double, double);

//###############
typedef enum {Livre=0,	// mensagem está livre para ser utilizada
	Oferta=1,	// oferta é postada pelo fornecedor
	   Proposta=2, // Comprador estabelece uma proposta de compra
	Aceite=3,	// Comprador aceita a proposta
	Recusa=4	// Comprador nao aceita a proposta
   } MsgStatus; 

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