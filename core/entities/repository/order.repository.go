package repository

import (
	"coding-chelleng/core/application/dtos"
	"coding-chelleng/core/entities/models"

	"gorm.io/gorm"
)

type OrderRepository interface {
	AddOrder(order models.Order) error
	DeleteOrder(id uint) error
	CountOrder(id uint) (int64, error)
	GetOrders() ([]models.Order, error)
	UpdateOrder(id uint, body dtos.UpdateRequestOrder) error
}
type OrderRepositoryImpl struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &OrderRepositoryImpl{db: db}
}
func (r *OrderRepositoryImpl) AddOrder(order models.Order) error {
	return r.db.Create(&order).Error
}
func (r *OrderRepositoryImpl) UpdateOrder(id uint, body dtos.UpdateRequestOrder) error {
	var order models.Order
	return r.db.Model(&order).Where("id = ?", id).Updates(&body).Error
}
func (r *OrderRepositoryImpl) DeleteOrder(id uint) error {
	var order models.Order
	return r.db.Where("id = ?", id).Delete(&order).Error
}
func (r *OrderRepositoryImpl) CountOrder(id uint) (int64, error) {
	var count int64
	var order models.Order
	err := r.db.Model(&order).Where("id = ?", id).Count(&count).Error
	return count, err
}
func (r *OrderRepositoryImpl) GetOrders() ([]models.Order, error) {
	var order []models.Order
	err := r.db.Find(&order).Error
	return order, err
}
