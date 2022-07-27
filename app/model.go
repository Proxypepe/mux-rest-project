package app

import (
	"database/sql"
	"errors"
	_ "github.com/lib/pq"
)

type product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func (p *product) getProduct(db *sql.DB) error {
	return errors.New("not implemented")
}

func (p *product) updateProduct(db *sql.DB) error {
	return errors.New("not implemented")
}

func (p *product) deleteProduct(db *sql.DB) error {
	return errors.New("not implemented")
}

func (p *product) createProduct(db *sql.DB) error {
	return errors.New("not implemented")
}

func (p *product) getProducts(db *sql.DB, start, count int) ([]product, error) {
	return nil, errors.New("not implemented")
}
