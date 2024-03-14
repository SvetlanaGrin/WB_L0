package repository

import (
	"WB_L0/internal/models"
	"context"
	_ "database/sql"
	"encoding/json"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/patrickmn/go-cache"
	"log"
)

type OrderPostgres struct {
	db *pgxpool.Pool
	c  *cache.Cache
}

type Order interface {
	GetOrderById(orderId string) (models.Order, error)
	AddOrder(order models.Order) error
}

func NewOrderPostgres(db *pgxpool.Pool, c *cache.Cache) *OrderPostgres {
	err := GetAllOrderInitCache(db, c)
	if err != nil {
		log.Fatalf("don't init cache: %v", err)
	}
	return &OrderPostgres{db: db, c: c}
}
func GetAllOrderInitCache(db *pgxpool.Pool, c *cache.Cache) error {
	row, err := db.Query(context.Background(), "SELECT * FROM orders")
	if err != nil {
		panic(err)
	}
	for row.Next() {
		order := make(map[string]interface{})
		var id string
		err = row.Scan(&id, &order)
		if err != nil {
			panic(err)
		}
		c.Set(id, order, cache.NoExpiration)
	}
	return nil
}

func (op *OrderPostgres) GetOrderById(orderId string) (models.Order, error) {
	order := models.Order{}
	orderJson, found := op.c.Get(orderId)
	if found {
		switch orderJson.(type) {
		case map[string]interface{}:
			b, err := json.Marshal(orderJson.(map[string]interface{}))
			if err != nil {
				log.Fatalf("failed Marshal order: %v", err)
				return order, err
			}
			err = json.Unmarshal(b, &order)
			if err != nil {
				log.Fatalf("failed Unmarshal order: %v", err)
				return order, err
			}
		case []uint8:
			err := json.Unmarshal(orderJson.([]uint8), &order)
			if err != nil {
				log.Fatalf("failed Unmarshal order: %v", err)
				return order, err
			}
		}

	}
	return order, nil
}

func (op *OrderPostgres) AddOrder(order models.Order) error {
	body, err := json.Marshal(&order)
	if err != nil {
		log.Fatalf("failed Marshal order: %v", err)
		return err
	}
	op.c.Set(order.OrderUid, body, cache.NoExpiration)
	_, found := op.c.Get(order.OrderUid)
	if found {
		_, err = op.db.Exec(context.Background(), `INSERT INTO orders ("uid", "order") VALUES($1,$2)`, order.OrderUid, body)
		if err != nil {
			panic(err)
		}
	}

	return nil
}
