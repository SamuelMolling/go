package comprador

import (
	"context"
	"fmt"
	"math/rand"
	quadromensagens "mercado_de_energia/pkg/quadro_mensagens"
	"time"
)

type EConsumidor struct {
	Id              int     //id
	TarifaDesejavel float64 // tarifa desejavel
	PrecoMaximo     float64 // maximo preco admissivel
	PrazoContrato   int     // prazo do contrato do cliente
	Demanda         float64 // demanda do cliente

	OfertaAberta bool                     // oferta aberta
	Oferta       *quadromensagens.MsgMerc // oferta
	Quadro       *quadromensagens.QuadroMsg
}

// Inicialização da estrutura de dados
func (c *EConsumidor) Inicia_EConsumidor() {
	c.OfertaAberta = false

	// fmt.Println("###########################")
	// fmt.Printf("Cadastrar dados Consumidor: %d\n", c.Id)
	// fmt.Println("###########################\n")
	// fmt.Print("Prazo de contrato do Consumidor [s]:")
	// valor := setValores()
	// c.PrazoContrato = int(valor)
	// fmt.Print("Demanda Consumidor [kW]:")
	// valor = setValores()
	// c.Demanda = valor
	// fmt.Print("Máximo preço admissível [R$/kW]:")
	// valor = setValores()
	// c.PrecoMaximo = valor
	// fmt.Print("Tarifa desejável [R$/kW]:")
	// valor = setValores()
	// c.TarifaDesejavel = valor

	rand.Seed(time.Now().UnixNano())
	c.PrazoContrato = rand.Intn(120)
	c.Demanda = rand.Float64() * 100
	c.PrecoMaximo = rand.Float64() * 100
	c.TarifaDesejavel = rand.Float64() * 100
}

func (c *EConsumidor) AtualizaPA() { //Atualiza preço máximo, caso o prazo esteja acabando
	if c.PrazoContrato == 15 { //se o prazo for menor ou igual que 5 segundos
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
	c.Quadro = &q

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	tick := time.NewTicker(time.Second * 1)
	terminoContrato := time.NewTimer(time.Second * time.Duration(c.PrazoContrato))

	for {
		select {
		case <-ctx.Done():
			return
		case <-tick.C:
			c.PrazoContrato--
			c.AtualizaPA()
		case <-terminoContrato.C:
			c.OfertaAberta = false
			q.MuRW.Lock()
			c.Oferta.Clean()
			q.MuRW.Unlock()
			return
		default:
			if c.PrazoContrato <= 0 || c.Demanda <= 0 {
				return
			}
			if c.Demanda > 0 && !c.OfertaAberta {
				oferta := &quadromensagens.MsgMerc{}   //Cria uma variável tipo quadro
				oferta.CodigoComprador = c.Id          //Vincula o id de um comprador
				oferta.DemandaSolicitada = c.Demanda   //Vincula uma demanda de um comprador
				oferta.Status = quadromensagens.Oferta //Vincula uma proposta de um comprador
				oferta.CodigoFornecedor = -1           //Vincula um fornecedor a uma proposta
				if index := q.SetQMsg(oferta); index == -1 {
					fmt.Println("\nQuadro de mensagens cheio")
					time.Sleep(time.Second * 5)
				} else {
					fmt.Printf("\nConsumidor %d enviou uma oferta de %.2f kW para o fornecedor (index %d) \n", c.Id, c.Demanda, index)
					c.OfertaAberta = true
					c.Oferta = oferta
				}
			}

			if c.OfertaAberta && c.Oferta.Status == quadromensagens.Proposta {
				if c.Oferta.PrecoVenda <= c.PrecoMaximo {
					fmt.Printf("\nConsumidor %d aceitou a oferta de %f kW de %f R$\n",
						c.Oferta.CodigoComprador,
						c.Oferta.CapacidadeFornecimento,
						c.Oferta.PrecoVenda)

					if c.Demanda < c.Oferta.CapacidadeFornecimento {
						c.Oferta.CapacidadeFornecimento = c.Demanda
					}

					c.Demanda -= c.Oferta.CapacidadeFornecimento
					c.Oferta.Status = quadromensagens.Aceite
					c.OfertaAberta = false
				} else {
					fmt.Printf("\nConsumidor %d recusou a oferta de %f kW de %f R$\n",
						c.Oferta.CodigoComprador,
						c.Oferta.CapacidadeFornecimento,
						c.Oferta.PrecoVenda)
					c.Oferta.Status = quadromensagens.Recusa
				}

			}
		}
	}
}

//Teste para print de threads
func PrintThreads(id int) {
	fmt.Printf("\nThread %d Execução em: ", id)
	printDate()
}
func printDate() { //Função para print data e hora atual
	currentTime := time.Now()
	fmt.Println(currentTime.Format("02/01/2006 15:04:05"))
}
