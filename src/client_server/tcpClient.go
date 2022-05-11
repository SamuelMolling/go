package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"

	"github.com/mackerelio/go-osstat/cpu"
	"github.com/mackerelio/go-osstat/memory"
)

func main() {

	for {

		conexao, err := net.Dial("tcp", "127.0.0.1:1234") //Monta a conexão
		if err != nil {
			fmt.Println(err)
			return
		}
		var texto [5]string
		for i := 0; i < 5; i++ {

			total_memory, used_memory := createMetricsMemory()   //Cria metricas
			user_cpu, system_cpu, idle_cpu := createMetricsCpu() //Cria metricas

			switch i {
			case 0:
				texto[0] = fmt.Sprintf("CPU User: %.2f%%\n", user_cpu)
				fmt.Fprintf(conexao, texto[0]+"\n")
				break
			case 1:
				texto[1] = fmt.Sprintf("CPU System: %.2f%%\n", system_cpu)
				fmt.Fprintf(conexao, texto[1]+"\n")
				break
			case 2:
				texto[2] = fmt.Sprintf("CPU Idle: %.2f%%\n", idle_cpu)
				fmt.Fprintf(conexao, texto[2]+"\n")
				break
			case 3:
				texto[3] = fmt.Sprintf("Memory Total: %dGB\n", total_memory/1000000000)
				fmt.Fprintf(conexao, texto[3]+"\n")
				break
			case 4:
				texto[4] = fmt.Sprintf("memory used: %dGB\n", used_memory/1000000000)
				fmt.Fprintf(conexao, texto[4]+"\n")
				break
			}

		}
		message, _ := bufio.NewReader(conexao).ReadString('\n')
		fmt.Print("Message send -> " + message)
	}
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

func createMetricsMemory() (uint64, uint64) { //Create metrics of Memory
	memory, err := memory.Get()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(-1)
	}
	total_memory := memory.Total
	used_memory := memory.Used
	return total_memory, used_memory
}
