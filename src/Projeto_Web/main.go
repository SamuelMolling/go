package main

import (
	"net/http"

	"github.com/SamuelMolling/routes"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
