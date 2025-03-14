package handlers

import (
	"context"
	"database/sql"
	"encoding/json"
	"github.com/Vinidamiaop/desafio-go-client-server-api/entities"
	"github.com/Vinidamiaop/desafio-go-client-server-api/utils"
	"io"
	"log"
	"net/http"
	"time"
)

func CotacaoHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()
		ctx, cancel := context.WithTimeout(ctx, 200*time.Millisecond)
		defer cancel()

		w.Header().Set("Content-Type", "application/json")

		req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			err = json.NewEncoder(w).Encode(utils.Response[any]{
				Message:   err.Error(),
				Data:      nil,
				IsSuccess: false,
			})
			return
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			err = json.NewEncoder(w).Encode(utils.Response[any]{
				Message:   err.Error(),
				Data:      nil,
				IsSuccess: false,
			})
			return
		}

		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			err = json.NewEncoder(w).Encode(utils.Response[any]{
				Message:   err.Error(),
				Data:      nil,
				IsSuccess: false,
			})
			return
		}

		var cotacao entities.Cotacao
		err = json.Unmarshal(body, &cotacao)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			err = json.NewEncoder(w).Encode(utils.Response[any]{
				Message:   err.Error(),
				Data:      nil,
				IsSuccess: false,
			})
			return
		}

		err = cotacao.Save(db)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			err = json.NewEncoder(w).Encode(utils.Response[any]{
				Message:   err.Error(),
				Data:      nil,
				IsSuccess: false,
			})
			return
		}

		err = json.NewEncoder(w).Encode(utils.Response[entities.Cotacao]{
			Message:   "Cotação salva com sucesso",
			Data:      cotacao,
			IsSuccess: true,
		})

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			err = json.NewEncoder(w).Encode(utils.Response[any]{
				Message:   err.Error(),
				Data:      nil,
				IsSuccess: false,
			})
			return
		}
	}
}
