package main

import (
	"github.com/Vinidamiaop/desafio-go-client-server-api/database"
	"github.com/Vinidamiaop/desafio-go-client-server-api/handlers"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
)

func main() {
	db, err := database.InitDB()
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.HealthcheckHandler)
	mux.HandleFunc("/cotacao", handlers.CotacaoHandler(db))

	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatalln(err)
	}
}
