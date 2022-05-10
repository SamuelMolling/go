package client

import (
	"fmt"
	"net"
	"os"
	"time"

	cpu "github.com/mackerelio/go-osstat/cpu"
	memory "github.com/mackerelio/go-osstat/memory"
)

func CreateClient() {

	conexao, erro1 := net.Dial("tcp", "127.0.0.1:50001") //Testa conexão
	if erro1 != nil {
		fmt.Println(erro1)
		os.Exit(3)
	}

	for {
		total_memory := createMetricsMemory() //Cria metricas

		total_memory_message := fmt.Sprintf("Memory Total: %dGB\n", total_memory/1000000000)
		fmt.Fprintf(conexao, total_memory_message) //Envia pela conexao TCP a mensagem

		user_cpu, system_cpu, idle_cpu := createMetricsCpu() //Cria metricas
		var user_cpu_message string = fmt.Sprintf("CPU User: %.2f%%\n", user_cpu)
		fmt.Fprintf(conexao, user_cpu_message) //Envia pela conexao TCP a mensagem

		var system_cpu_message string = fmt.Sprintf("CPU System: %.2f%%\n", system_cpu)
		fmt.Fprintf(conexao, system_cpu_message) //Envia pela conexao TCP a mensagem

		var idle_cpu_message string = fmt.Sprintf("CPU Idle: %.2f%%\n", idle_cpu)
		fmt.Fprintf(conexao, idle_cpu_message) //Envia pela conexao TCP a mensagem
	}

}

func createMetricsMemory() uint64 { //Create metrics of Memory
	memory, err := memory.Get()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(-1)
	}
	total_memory := memory.Total
	return total_memory
}

func createMetricsCpu() (float64, float64, float64) { //Create metrics of Cpu
	before, err := cpu.Get()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
	}
	time.Sleep(time.Duration(1) * time.Second)
	after, err := cpu.Get()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
	}
	total := float64(after.Total - before.Total)
	user_cpu := float64(after.User-before.User) / total * 100
	system_cpu := float64(after.System-before.System) / total * 10
	idle_cpu := float64(after.Idle-before.Idle) / total * 100
	return user_cpu, system_cpu, idle_cpu
}
