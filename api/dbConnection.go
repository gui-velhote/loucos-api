package api

import (
	"context"
	"log"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/jackc/pgx/v5/pgxpool"
)

var config Config

type Config struct {
	Database struct {
		USER                   string `toml:"user"`
		PASSWORD               string `toml:"password"`
		NET                    string `toml:"net"`
		ADDRESS                string `toml:"address"`
		DATABASE_NAME          string `toml:"database_name"`
		ALLOW_NATIVE_PASSWORDS bool   `toml:"allow_native_passwords"`
	}
}

func parseDatbaseConfig(FILE_PATH string) Config {

	var config Config

	data, err := os.ReadFile(FILE_PATH)

	if err != nil {
		log.Fatal(err)
	}

	err = toml.Unmarshal(data, &config)

	if err != nil {
		log.Fatal(err)
	}

	return config
}

func databaseConnection() *pgxpool.Pool {
  
  dbpool, err := pgxpool.New(context.Background(), "user=loki password=2002 dbname=dragnarok sslmode=verify-ca pool_max_conns=10")

  if err != nil {
    log.Fatal("Erro ao obter conexão com banco de dados")
    os.Exit(1)
  }

	return dbpool
}
