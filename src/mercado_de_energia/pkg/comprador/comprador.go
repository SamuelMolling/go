package comprador

import (
	"context"
	"fmt"
	quadromensagens "mercado_de_energia/pkg/quadro_mensagens"
)

type EConsumidor struct {
	Id              int     //id
	TarifaDesejavel float64 // tarifa desejavel
	PrecoMaximo     float64 // maximo preco admissivel
	PrazoContrato   float64 // prazo do contrato do cliente
	Demanda         float64 // demanda do cliente
}

// Inicialização da estrutura de dados
func (c *EConsumidor) Inicia_EConsumidor() {
	fmt.Printf("\nCadastrar dados Consumidor: %d\n", c.Id)
	fmt.Print("Prazo de contrato do Consumidor [s]:")
	valor := setValores()
	c.PrazoContrato = valor
	fmt.Print("Demanda Consumidor [kW]:")
	valor = setValores()
	c.Demanda = valor
	fmt.Print("Máximo preço admissível [R$/kW]:")
	valor = setValores()
	c.PrecoMaximo = valor
	fmt.Print("Tarifa desejável [R$/kW]:")
	valor = setValores()
	c.TarifaDesejavel = valor
}

func (c *EConsumidor) AtualizaPA() { //Atualiza preço máximo, caso o prazo esteja acabando
	if c.PrazoContrato <= 5 {
		c.PrecoMaximo += (c.PrecoMaximo * 0.22)
	} else {
		c.TarifaDesejavel += 0.05
	}
}

func setValores() float64 {
	var valor float64
	fmt.Scan(&valor)
	fmt.Println("")
	return valor
}

func (c *EConsumidor) AtualizaDemanda(demanda_contratada float64) float64 { //Atualiza a demanda geral quando uma demanda for contratada
	c.Demanda -= demanda_contratada
	return c.Demanda
}

func (c *EConsumidor) WorkConsumidor(ctx context.Context, q quadromensagens.QuadroMsg) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			quadro := quadromensagens.MsgMerc{}      //Cria uma variável tipo quadro
			quadro.CodigoComprador = c.Id            //Vincula o id de um comprador
			quadro.DemandaSolicitada = c.Demanda     //Vincula uma demanda de um comprador
			quadro.Status = quadromensagens.Proposta //Vincula uma proposta de um comprador
			q.MuRW.Lock()                            //Cria um lock mutex de Read e Write
			//q.Mensagem = append(q.Mensagem[:q.LivreQMsg()], q.Mensagem[q.ProxQMsg():]...) //Elimina a mensagem atual caso ela esteja livre
			if len(q.Mensagem) < 8 { //valida se o tamanho do array é menor que 8
				q.Mensagem = append(q.Mensagem, quadro) //Cria a mensagem, caso seja menor que 8
			}
			q.MuRW.Unlock() //Desbloqueia o Mutex
		}
	}
}
