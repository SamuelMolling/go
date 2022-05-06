package comprador

type EConsumidor struct {
	Id              int     //id
	TarifaDesejavel float64 // tarifa desejavel
	PrecoMaximo     float64 // maximo preco admissivel
	PrazoContrato   float64 // prazo do contrato do cliente
	Demanda         float64 // demanda do cliente
}

// Inicialização da estrutura de dados
func (c *EConsumidor) Inicia_EConsumidor() *EConsumidor {

	// rand.Seed(time.Now().Unix())
	// c.TarifaDesejavel = fornecedor.GetRandFloat(500, 700)
	// c.PrecoMaximo = fornecedor.GetRandFloat(500, 800)
	// c.PrazoContrato = fornecedor.GetRandFloat(1, 4)
	// c.Demanda = fornecedor.GetRandFloat(500, 1000)
	return c
}

func (c *EConsumidor) AtualizaPA() {

}

//Atualiza tA
//void atualiza_pa(struct_EConsumidor *, double, double, double)
