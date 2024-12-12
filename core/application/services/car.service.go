package services

import (
	"coding-chelleng/core/application/dtos"
	"coding-chelleng/core/entities/models"
	"coding-chelleng/core/entities/repository"
	"coding-chelleng/pkg"
	"strconv"

	"github.com/go-playground/validator/v10"
)

type CarService interface {
	AddCar(name string, dRate float32, mRate float32, img string) error
	GetCarByID(id string) (*models.Car, error)
	UpdateCar(id string, body dtos.UpdateRequestCar) error
	DeleteCar(id string) error
	GetCars() ([]models.Car, error)
}

type CarServiceImpl struct {
	repo     repository.CarRepository
	validate *validator.Validate
}

func NewCarService(repo repository.CarRepository, validate *validator.Validate) CarService {
	return &CarServiceImpl{repo: repo, validate: validate}
}
func (s *CarServiceImpl) AddCar(name string, dRate float32, mRate float32, img string) error {

	car := models.Car{
		Name:      name,
		DayRate:   dRate,
		MonthRate: mRate,
		Image:     img,
	}
	if err := s.repo.AddCar(car); err != nil {
		return err
	}
	return nil
}
func (s *CarServiceImpl) GetCarByID(id string) (*models.Car, error) {
	carId, _ := strconv.ParseUint(id, 10, 32)
	count, _ := s.repo.CountCarByID(uint(carId))
	if count == 0 {
		return nil, &pkg.NotFound{Message: "car not found"}
	}
	car, err := s.repo.GetCarByID(uint(carId))
	if err != nil {
		return nil, err
	}
	return car, nil
}
func (s *CarServiceImpl) UpdateCar(id string, body dtos.UpdateRequestCar) error {
	if err := s.validate.Struct(body); err != nil {
		return &pkg.BadRequest{Message: err.Error()}
	}
	carId, _ := strconv.ParseUint(id, 10, 32)
	count, _ := s.repo.CountCarByID(uint(carId))
	if count == 0 {
		return &pkg.NotFound{Message: "car not found"}
	}
	if err := s.repo.UpdateCar(uint(carId), body); err != nil {
		return err
	}
	return nil
}
func (s *CarServiceImpl) DeleteCar(id string) error {
	carId, _ := strconv.ParseUint(id, 10, 32)
	count, _ := s.repo.CountCarByID(uint(carId))
	if count == 0 {
		return &pkg.NotFound{Message: "car not found"}
	}
	if err := s.repo.DeleteCar(uint(carId)); err != nil {
		return err
	}
	return nil
}
func (s *CarServiceImpl) GetCars() ([]models.Car, error) {
	cars, err := s.repo.GetCars()
	if err != nil {
		return nil, err
	}
	return cars, err
}
