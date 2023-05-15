package postgresql

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
	"github.com/mahiro72/go_api-template/pkg/config"
)

func NewDB() (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		config.Env.POSTGRES_USER,
		config.Env.POSTGRES_PASSWORD,
		config.Env.POSTGRES_HOST,
		"5432",
		config.Env.POSTGRES_DB,
	)
	db, err := sql.Open("postgres", dsn)

	if err != nil {
		tryConnCnt := 1
		for tryConnCnt <= 10 && err != nil {
			timer := time.NewTimer(2 * time.Second)
			db, err = sql.Open("postgres", dsn)
			<-timer.C

			log.Println("NewDB: tryConnDB =", tryConnCnt)
			tryConnCnt += 1
		}
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
