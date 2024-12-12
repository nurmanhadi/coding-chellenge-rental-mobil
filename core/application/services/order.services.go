package services

import (
	"coding-chelleng/core/application/dtos"
	"coding-chelleng/core/entities/models"
	"strconv"
	"time"

	"coding-chelleng/core/entities/repository"
	"coding-chelleng/pkg"

	"github.com/go-playground/validator/v10"
)

type OrderService interface {
	AddOrder(body dtos.AddRequestOrder) error
	DeleteOrder(id string) error
	GetOrders() ([]models.Order, error)
	UpdateOrder(id string, body dtos.UpdateRequestOrder) error
}
type OrderServiceImpl struct {
	repo     repository.OrderRepository
	validate *validator.Validate
}

func NewOrderService(repo repository.OrderRepository, validate *validator.Validate) OrderService {
	return &OrderServiceImpl{repo: repo, validate: validate}
}

func (s *OrderServiceImpl) AddOrder(body dtos.AddRequestOrder) error {
	if err := s.validate.Struct(body); err != nil {
		return &pkg.BadRequest{Message: err.Error()}
	}
	orderDate, _ := time.Parse("2006-01-02", body.OrderDate)
	pickupDate, _ := time.Parse("2006-01-02", body.PickupDate)
	dropoffDate, _ := time.Parse("2006-01-02", body.DropOffDate)
	order := models.Order{
		CarId:           body.CarId,
		OrderDate:       orderDate,
		PickupDate:      pickupDate,
		DropOffDate:     dropoffDate,
		PickupLocation:  body.PickupLocation,
		DropOffLocation: body.DropOffLocation,
	}
	if err := s.repo.AddOrder(order); err != nil {
		return err
	}
	return nil
}
func (s *OrderServiceImpl) UpdateOrder(id string, body dtos.UpdateRequestOrder) error {
	if err := s.validate.Struct(body); err != nil {
		return &pkg.BadRequest{Message: err.Error()}
	}
	orderId, _ := strconv.ParseUint(id, 10, 32)
	count, _ := s.repo.CountOrder(uint(orderId))
	if count == 0 {
		return &pkg.NotFound{Message: "car not found"}
	}
	if err := s.repo.UpdateOrder(uint(orderId), body); err != nil {
		return err
	}
	return nil
}
func (s *OrderServiceImpl) DeleteOrder(id string) error {
	orderId, _ := strconv.ParseUint(id, 10, 32)
	count, _ := s.repo.CountOrder(uint(orderId))
	if count == 0 {
		return &pkg.NotFound{Message: "car not found"}
	}

	if err := s.repo.DeleteOrder(uint(orderId)); err != nil {
		return err
	}
	return nil
}
func (s *OrderServiceImpl) GetOrders() ([]models.Order, error) {
	orders, err := s.repo.GetOrders()
	if err != nil {
		return nil, err
	}
	return orders, nil
}
