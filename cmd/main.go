package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

var config struct {
	dsn string
}

func main() {
	flag.StringVar(&config.dsn, "db-dsn", "postgresql://go-postgres:pa55word@localhost/go-postgres", "PostgreSQL DSN")
	db, err := openDB(config.dsn)
	if err != nil {
		log.Fatal(err)
	}

	shop := NewFruitShop(db)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	// insert
	banana := &Fruit{Name: "Banana", Origin: "Taiwan", Color: "yellow", Created_at: time.Now()}
	err = shop.Insert(ctx, banana)
	if err != nil {
		log.Fatal("fail to insert fruit: %v", err)
	}

	// get by id
	// get by title
	// full text search
	// update
	// partial update
	// delete
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("fail to open database: %w", err)
	}
	return db, nil
}
