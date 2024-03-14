package service

import (
	"WB_L0/internal/models"
	"WB_L0/internal/repository"
)

type OrderService struct {
	repo repository.Order
}

type Order interface {
	GetOrderById(orderId string) (models.Order, error)
	AddOrder(order models.Order) error
	//UpdatePlus(userid int, input AvitoTest.UserBalance) error
	//UpdateMinus(userid int, input AvitoTest.UserBalance) error
}

func NewOrderService(repo repository.Order) *OrderService {
	return &OrderService{repo: repo}
}

func (oSer *OrderService) GetOrderById(orderId string) (models.Order, error) {
	return oSer.repo.GetOrderById(orderId)
}

func (oSer *OrderService) AddOrder(order models.Order) error {
	return oSer.repo.AddOrder(order)
}
