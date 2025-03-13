package database

import (
	"context"
	"database/sql"
	"time"
)

type Cotacao struct {
	Usdbrl struct {
		Code       string `json:"code"`
		Codein     string `json:"codein"`
		Name       string `json:"name"`
		High       string `json:"high"`
		Low        string `json:"low"`
		VarBid     string `json:"varBid"`
		PctChange  string `json:"pctChange"`
		Bid        string `json:"bid"`
		Ask        string `json:"ask"`
		Timestamp  string `json:"timestamp"`
		CreateDate string `json:"create_date"`
	} `json:"USDBRL"`
}

func (c *Cotacao) Save(db *sql.DB) error {
	ctx := context.Background()
	ctx, dbCancel := context.WithTimeout(ctx, 10*time.Millisecond)
	defer dbCancel()

	stmt, err := db.Prepare("insert into cotacao (code, codein, name, high, low, varBid, pctChange, bid, ask, timestamp, create_date) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, c.Usdbrl.Code, c.Usdbrl.Codein, c.Usdbrl.Name, c.Usdbrl.High, c.Usdbrl.Low, c.Usdbrl.VarBid, c.Usdbrl.PctChange, c.Usdbrl.Bid, c.Usdbrl.Ask, c.Usdbrl.Timestamp, c.Usdbrl.CreateDate)
	if err != nil {
		return err
	}

	return nil
}
