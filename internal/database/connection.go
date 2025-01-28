package database

import (
	"context"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

var Conn *pgxpool.Pool

// Função para estabelecer conexão com o bando de dados
func NewConnection(connectionString string) (*pgxpool.Pool, error) {
	//Contexto criado para controlar a execução de operações assíncronas
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	//Pool de conexões para indicar se foi bem-sucedida
	var err error
	Conn, err = pgxpool.Connect(ctx, connectionString)
	if err != nil {
		return nil, err
	}

	return Conn, nil
}
