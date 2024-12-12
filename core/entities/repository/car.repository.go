package repository

import (
	"coding-chelleng/core/application/dtos"
	"coding-chelleng/core/entities/models"

	"gorm.io/gorm"
)

type CarRepository interface {
	AddCar(car models.Car) error
	GetCarByID(id uint) (*models.Car, error)
	CountCarByID(id uint) (int64, error)
	UpdateCar(id uint, body dtos.UpdateRequestCar) error
	DeleteCar(id uint) error
	GetCars() ([]models.Car, error)
}

type CarRepositoryImpl struct {
	db *gorm.DB
}

func NewCarRepository(db *gorm.DB) CarRepository {
	return &CarRepositoryImpl{db: db}
}
func (r *CarRepositoryImpl) AddCar(car models.Car) error {
	return r.db.Create(&car).Error
}
func (r *CarRepositoryImpl) GetCarByID(id uint) (*models.Car, error) {
	var car models.Car
	err := r.db.Where("id = ?", id).First(&car).Error
	return &car, err
}
func (r *CarRepositoryImpl) CountCarByID(id uint) (int64, error) {
	var count int64
	var car models.Car
	err := r.db.Model(&car).Where("id = ?", id).Count(&count).Error
	return count, err
}
func (r *CarRepositoryImpl) UpdateCar(id uint, body dtos.UpdateRequestCar) error {
	var car models.Car
	return r.db.Model(&car).Where("id = ?", id).Updates(&body).Error
}
func (r *CarRepositoryImpl) DeleteCar(id uint) error {
	var car models.Car
	return r.db.Where("id = ?", id).Delete(&car).Error
}
func (r *CarRepositoryImpl) GetCars() ([]models.Car, error) {
	var car []models.Car
	err := r.db.Find(&car).Error
	return car, err
}
