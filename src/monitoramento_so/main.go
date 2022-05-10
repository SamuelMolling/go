package main

import (
	"context"
	"fmt"
	"monitoramento_so/pkg/client"
	"monitoramento_so/pkg/server"
	"time"
)

func main() {

	ctx, err := context.WithTimeout(context.Background(), 20*time.Second)
	if err == nil {
		fmt.Println(err)
	} else {
		go server.CreateServer()
		go client.CreateClient()
		<-ctx.Done()
	}
}
