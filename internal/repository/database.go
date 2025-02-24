package repository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/evrintobing17/my-superindo-app/config"
	_ "github.com/lib/pq"
)

type Database struct {
	Conn *sql.DB
}

func NewDatabase(cfg config.DatabaseConfig) (*Database, error) {
	Dbdriver := "postgres"
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.Username, cfg.Name)
	db, err := sql.Open(Dbdriver, dsn)
	if err != nil {
		fmt.Println("Cannot connect to database")
		log.Fatal("Error: ", err)
		return nil, err
	}

	var dbName string
	err = db.QueryRow("SELECT current_database()").Scan(&dbName)
	if err != nil {
		log.Fatal("Failed to fetch current database:", err)
	}
	log.Println("Connected to database:", dbName)
	if dbName != cfg.Name {
		panic("hahahahah")
	}
	log.Printf("We are connected to the %s database\n", Dbdriver)

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return &Database{Conn: db}, nil
}

func (db *Database) Close() error {
	return db.Conn.Close()
}
