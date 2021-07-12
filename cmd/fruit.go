package main

import (
	//"fmt"
	"context"
	"database/sql"
	"log"
)

type FruitShop struct {
	DB *sql.DB
}

func NewFruitShop(db *sql.DB) *FruitShop {
	return &FruitShop{DB: db}
}

func (f *FruitShop) GetByID(ctx context.Context, id int64) error {
	return nil
}

func (f *FruitShop) Insert(ctx context.Context, fruit *Fruit) error {
	log.Printf("%+v inserted", fruit)
	return nil
}

func (f *FruitShop) Update(ctx context.Context, fruit *Fruit) (*Fruit, error) {
	log.Printf("%+v updated", fruit)
	return fruit, nil
}

func (f *FruitShop) Delete(ctx context.Context, id int64) error {
	log.Printf("id: %d deleted", id)
	return nil
}
