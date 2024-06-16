package postges

import (
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx"
	"github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

// func NewGenerated() {
// 	dbHost := os.Getenv("DB_HOST")
// 	dbPort := os.Getenv("DB_PORT")
// 	dbUser := os.Getenv("DB_USER")
// 	dbPassword := os.Getenv("DB_PASSWORD")
// 	dbName := os.Getenv("DB_NAME")
//
// 	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
// 		dbHost, dbPort, dbUser, dbPassword, dbName)
//
// 	db, err := sql.Open("postgres", psqlInfo)
// 	if err != nil {
// 		log.Fatalf("Failed to connect to database: %v\n", err)
// 	}
// 	defer db.Close()
//
// 	err = db.Ping()
// 	if err != nil {
// 		log.Fatalf("Failed to ping database: %v\n", err)
// 	}
//
// 	log.Println("Successfully connected to the database!")
// 	// Your application code here...
// }

func NewDB() (*sqlx.DB, error) {
	db, err := connect()
	if err != nil {
		return nil, fmt.Errorf("connect: %w", err)
	}

	err = pingDB(db)
	if err != nil {
		return nil, fmt.Errorf("pingDB: %w", err)
	}

	return db, nil
}

func pingDB(db *sqlx.DB) error {
	if db == nil {
		return errors.New("db is nil")
	}

	if err := db.Ping(); err == nil {
		return nil
	}

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	timer := time.NewTimer(30 * time.Second)
	defer timer.Stop()

	for {
		select {
		case <-ticker.C:
			if err := db.Ping(); err == nil {
				return nil
			}
		case <-timer.C:
			fmt.Println("Tick db overtime")
			return db.Ping()
		}
	}
}

func connect() (*sqlx.DB, error) {
	pgxConfig, err := pgx.ParseEnvLibpq()
	if err != nil {
		return nil, fmt.Errorf("can't get pgx config: %w", err)
	}

	sqlDB := stdlib.OpenDB(pgxConfig)

	return sqlx.NewDb(sqlDB, "pgx"), nil
}
