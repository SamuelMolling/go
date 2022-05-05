// Criando um tipo de dados que será utilizado pelas funções  de calculo da equacao linear
type Efornecedor struct {
		PrecoDesejavel		float64 // preco desejavel
		MenorPreco 			float64 // menor preco admissivel
		CapacidadeCarga 	float64	// capacidade de carga
 		EnergiaGerada 		float64 // energia gerada
 		Energia_Fornecida 	float64 // energia fornecida
 		Demanda_Interna 	float64	// demanda interna
	 	FazOferta			int //variavel que indica se fornecedor fez uma oferta (se ==1)
}



// Inicialização da estrutura de dados
void inicia_struct_Efornecedor(struct_Efornecedor *, int)

// Atualizacao do pD
void atualiza_pd(struct_Efornecedor *, double, double, double)

//Gerador de num aleatorios
double num_aleat (double, double)








