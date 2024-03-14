package service

import "WB_L0/internal/repository"

type Service struct {
	Order
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Order: NewOrderService(repos.Order),
	}
}
