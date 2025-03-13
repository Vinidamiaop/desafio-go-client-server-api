package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	getCotacao()
}

type Cotacao struct {
	Bid string `json:"bid"`
}

func getCotacao() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 300*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost:8080/cotacao", nil)
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var cotacao Cotacao
	err = json.Unmarshal(body, &cotacao)
	if err != nil {
		log.Fatalln(err)
	}

	err = SalvaCotacao(cotacao.Bid)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Cotação salva com sucesso")
}

func SalvaCotacao(cotacao string) error {
	file, erro := os.Create("cotacao.txt")
	if erro != nil {
		return erro
	}

	_, erro = file.WriteString(fmt.Sprintf("Dólar: %s", cotacao))
	if erro != nil {
		return erro
	}

	return nil
}
