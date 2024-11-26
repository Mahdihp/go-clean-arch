package db

import (
	"context"
	"entgo.io/ent/dialect"
	"fmt"
	"github.com/bxcodec/go-clean-arch/config"
	"github.com/bxcodec/go-clean-arch/domain/ent"
	_ "github.com/lib/pq"
	"log"
)

type PostgresDB struct {
	db     *ent.Client
	config config.Config
}

func NewPostgres(cfg config.Config) *PostgresDB {
	ConnetionString := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		cfg.Postgres.PostgresHost, cfg.Postgres.PostgresPort, cfg.Postgres.PostgresUser, cfg.Postgres.PostgresDB, cfg.Postgres.PostgresPassword)

	fmt.Println("ConnetionString: ", ConnetionString)
	client, err := ent.Open(dialect.Postgres, ConnetionString)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	return &PostgresDB{config: cfg, db: client}
}
func (m *PostgresDB) Conn() *ent.Client {
	return m.db
}
