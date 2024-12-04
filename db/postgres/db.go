package db

import (
	"context"
	"entgo.io/ent/dialect"
	"fmt"
	"github.com/bxcodec/go-clean-arch/config"
	"github.com/bxcodec/go-clean-arch/internal/bybit_history_service/models/ent"
	_ "github.com/lib/pq"
	"log"
)

type PostgresDB struct {
	db     *ent.Client
	config config.Postgres
}

func NewPostgres(cfg config.Postgres) *PostgresDB {
	ConnetionString := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		cfg.PostgresHost, cfg.PostgresPort, cfg.PostgresUser, cfg.PostgresDB, cfg.PostgresPassword)

	fmt.Println("ConnetionString: ", ConnetionString)
	client, err := ent.Open(dialect.Postgres, ConnetionString)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	return &PostgresDB{config: cfg, db: client}
}
func (m *PostgresDB) Conn() *ent.Client {
	return m.db
}
