package order

import (
	"submission-2/core"
	"submission-2/repo/order"
)

type Service struct {
	orderRepo *order.Repository
}

func NewService(orderRepo *order.Repository) *Service {
	return &Service{
		orderRepo: orderRepo,
	}
}

func (s *Service) CreateOrder(order core.Order) (core.Order, error) {
	res, err := s.orderRepo.CreateOrder(order)

	return res, err

}

func (s *Service) GetAllOrders() ([]core.Order, error) {
	res, err := s.orderRepo.GetAllOrders()

	return res, err
}

func (s *Service) UpdateOrder(order core.Order, id uint64) (core.Order, error) {
	res, err := s.orderRepo.UpdateOrder(order, id)

	return res, err
}

func (s *Service) DeleteOrder(id uint64) (string, error) {
	message, err := s.orderRepo.DeleteOrder(id)

	return message, err

}
