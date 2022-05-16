package comprador

import (
	"context"
	"fmt"
	"math/rand"
	quadromensagens "mercado_de_energia/pkg/quadro_mensagens"
	"time"
)

type EConsumidor struct {
	Id              int                        // id
	TarifaDesejavel float64                    // tarifa desejavel
	PrecoMaximo     float64                    // maximo preco admissivel
	PrazoContrato   int                        // prazo do contrato do cliente
	Demanda         float64                    // demanda do cliente
	OfertaAberta    bool                       // oferta aberta
	Oferta          *quadromensagens.MsgMerc   // oferta
	Quadro          *quadromensagens.QuadroMsg // quadro
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
		c.TarifaDesejavel += c.PrecoMaximo
		c.PrecoMaximo = c.TarifaDesejavel
	} else {
		c.TarifaDesejavel += 1 //Acrescenta 1 no valor da tarifa desejável
	}
}

func setValores() float64 { //Setar valores dos Econs
	var valor float64
	fmt.Scan(&valor)
	fmt.Println("")
	return valor
}

func (c *EConsumidor) AtualizaDemanda(demanda_contratada float64) float64 { //Atualiza a demanda geral quando uma demanda for contratada
	c.Demanda -= demanda_contratada
	return c.Demanda
}

func (c *EConsumidor) WorkConsumidor(ctx context.Context, q quadromensagens.QuadroMsg) { //Criação de um worker
	c.Quadro = &q

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	tick := time.NewTicker(time.Second * 1)
	terminoContrato := time.NewTimer(time.Second * time.Duration(c.PrazoContrato))

	for {
		select {
		case <-ctx.Done(): //Espera o timeout de 120
			return
		case <-tick.C:
			c.PrazoContrato-- //decrementa o prazo de contrato
			c.AtualizaPA()    //Atualiza preço desejável
			continue
		case <-terminoContrato.C:
			c.OfertaAberta = false //fecha oferta
			q.MuRW.Lock()          //lock quadro
			c.Oferta.Clean()       //limpa o quadro
			q.MuRW.Unlock()        //unlock
			return
		default:
			if c.PrazoContrato <= 0 || c.Demanda <= 0 {
				return
			}
			if c.Demanda > 0 && !c.OfertaAberta { //se demanda > 0 e oferta aberta seja false -> faz oferta
				q.MuRW.Lock()
				oferta := &quadromensagens.MsgMerc{}   //Cria uma variável tipo quadro
				oferta.CodigoComprador = c.Id          //Vincula o id de um comprador
				oferta.DemandaSolicitada = c.Demanda   //Vincula uma demanda de um comprador
				oferta.Status = quadromensagens.Oferta //Vincula uma proposta de um comprador
				oferta.CodigoFornecedor = -1           //Vincula um fornecedor a uma proposta
				q.MuRW.Unlock()
				if index := q.SetQMsg(oferta); index == -1 { //Valida se o indice é menor que
					time.Sleep(time.Second * 5) //sleep
				} else {
					fmt.Printf("\nConsumidor %d enviou uma proposta de %.2f kW para o fornecedor (quadro %d) \n", c.Id, c.Demanda, index)
					c.OfertaAberta = true
					c.Oferta = oferta
				}
			}

			if c.OfertaAberta && c.Oferta.Status == quadromensagens.Proposta { //se tiver proposta aberta enviada pelo fornecedor
				if c.Oferta.PrecoVenda <= c.TarifaDesejavel { //Valida se o preco de venda é menor que o preco maximo
					if c.Demanda < c.Oferta.CapacidadeFornecimento { //valida se a demanda é menor que a capacidade de fornecimento total do fornecedor
						c.Oferta.CapacidadeFornecimento = c.Demanda //Se for, coloca como a capacidade de fornecimento é o valor total da demanda (vai comprar tudo)
					}
					q.MuRW.Lock()
					fmt.Printf("\nConsumidor %d aceitou a oferta de %.2f kW por %.2f kW/R$\n", //aceitou
						c.Oferta.CodigoComprador,
						c.Oferta.CapacidadeFornecimento,
						c.Oferta.PrecoVenda)

					c.Demanda -= c.Oferta.CapacidadeFornecimento //Subtrai a demanda total, menos o que comprou
					c.Oferta.Status = quadromensagens.Aceite     //coloca a mensagem como aceite
					c.OfertaAberta = false                       //e seta que essa proposta não está aberta
					q.MuRW.Unlock()
				} else {
					fmt.Printf("\nConsumidor %d recusou a oferta de %.2f kW por %.2f kW/R$\n", //se preco for maior que o maximo, recusa a proposta
						c.Oferta.CodigoComprador,
						c.Oferta.CapacidadeFornecimento,
						c.Oferta.PrecoVenda)
					q.MuRW.Lock()
					c.Oferta.Status = quadromensagens.Recusa
					q.MuRW.Unlock()
				}
			}
		}
	}
}
