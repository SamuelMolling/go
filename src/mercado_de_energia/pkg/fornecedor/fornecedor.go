package fornecedor

import (
	"crypto/rand"
	"math/big"
	"fmt"
)
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




//Gerador de num aleatorios float
func num_aleat(min, max float64) float64 {
	const floatPrecision = 1000000 //precisao do calculo
	minInt := int(min * floatPrecision) 
	maxInt := int(max * floatPrecision)

	return float64(GetRandInt(minInt, maxInt)) / floatPrecision
}

//Gerador de num aleatorios int
func GetRandInt(min, max int) int {
	nBig, _ := rand.Int(rand.Reader, big.NewInt(int64(max+1-min)))
	n := nBig.Int64()
	return int(n) + min
}







