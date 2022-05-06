package comprador

type EConsumidor struct {
		TarifaDesejavel		float64 // tarifa desejavel
		PrecoMaximo  		float64	// maximo preco admissivel
		PrazoContrato		float64 // prazo do contrato do cliente
		Demanda 			float64	// demanda do cliente
} 

// Inicialização da estrutura de dados
void inicia_struct_EConsumidor(struct_EConsumidor *, int)


//Atualiza tA
void atualiza_pa(struct_EConsumidor *, double, double, double)
