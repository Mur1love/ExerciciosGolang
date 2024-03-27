package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func horaCerta(w http.ResponseWriter, r *http.Request) {
	data := time.Now().Format("02/01/2006")
	hora := time.Now().Format("03:04:05")
	fmt.Fprintf(w,
		`<h1>Data: %s<h1>
		<h1>Hora: %s<h1>`, data, hora)
}

func main() {
	http.HandleFunc("/horaCerta", horaCerta)
	log.Println("Executando...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
