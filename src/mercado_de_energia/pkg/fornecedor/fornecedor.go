package fornecedor

import (
	"math/rand"
	"time"
)

// Criando um tipo de dados que será utilizado pelas funções  de calculo da equacao linear
type Efornecedor struct {
	Id                int     // id
	PrecoDesejavel    float64 // preco desejavel
	MenorPreco        float64 // menor preco admissivel
	CapacidadeCarga   float64 // capacidade de carga
	EnergiaGerada     float64 // energia gerada
	Energia_Fornecida float64 // energia fornecida
	Demanda_Interna   float64 // demanda interna
	FazOferta         int     //variavel que indica se fornecedor fez uma oferta (se ==1)
}

// Inicialização da estrutura de dados
func (c *Efornecedor) Inicia_Efornecedor() *Efornecedor {

	rand.Seed(time.Now().Unix())
	c.PrecoDesejavel = GetRandFloat(500, 1000)
	c.MenorPreco = GetRandFloat(100, 500)
	c.CapacidadeCarga = GetRandFloat(100, 200)
	c.EnergiaGerada = GetRandFloat(5000, 10000)
	//c.Energia_Fornecida = c.EnergiaGerada - c.
	c.Demanda_Interna = GetRandFloat(0, 1000)
	c.FazOferta = 0
	return c
}

// Atualizacao do pD
//void atualiza_pd(struct_Efornecedor *, double, double, double)
func (c *Efornecedor) AtualizaPrecoDesejavel() float64 {
	c.PrecoDesejavel = GetRandFloat(500, 1000)
	return c.PrecoDesejavel
}

//Gerador de num aleatorios float
func GetRandFloat(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}
