package main

import (
	"context"
	"monitoramento_so/pkg/client"
	"monitoramento_so/pkg/server"
	"time"
)

func main() {

	ctx, _ := context.WithTimeout(context.Background(), 20*time.Second)
	go server.CreateServer()
	go client.CreateClient()
	<-ctx.Done()
}
