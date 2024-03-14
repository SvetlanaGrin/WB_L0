package repository

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/patrickmn/go-cache"
)

type Repository struct {
	Order
}

func NewRepository(db *pgxpool.Pool, c *cache.Cache) *Repository {
	return &Repository{
		Order: NewOrderPostgres(db, c),
	}
}
