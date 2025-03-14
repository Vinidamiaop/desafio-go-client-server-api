package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Vinidamiaop/desafio-go-client-server-api/entities"
	"github.com/Vinidamiaop/desafio-go-client-server-api/utils"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	getCotacao()
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

	var response utils.Response[entities.Cotacao]
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Fatalln(err)
	}

	if response.IsSuccess == false {
		log.Fatalln(response.Message)
	}

	err = SalvaCotacao(response.Data)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Cotação salva com sucesso")
}

func SalvaCotacao(cotacao entities.Cotacao) error {
	file, erro := os.Create("cotacao.txt")
	if erro != nil {
		return erro
	}

	_, erro = file.WriteString(fmt.Sprintf("Dólar: %s", cotacao.Usdbrl.Bid))
	if erro != nil {
		return erro
	}

	return nil
}
