package handlers

import (
	"coding-chelleng/core/application/dtos"
	"coding-chelleng/core/application/services"
	"coding-chelleng/pkg"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

type CarHandler interface {
	AddCar(c *fiber.Ctx) error
	GetCarById(c *fiber.Ctx) error
	UpdateCar(c *fiber.Ctx) error
	DeleteCar(c *fiber.Ctx) error
	GetCars(c *fiber.Ctx) error
}

type CarHandlerImpl struct {
	service services.CarService
}

func NewCarHandler(service *services.CarService) CarHandler {
	return &CarHandlerImpl{service: *service}
}
func (h *CarHandlerImpl) AddCar(c *fiber.Ctx) error {
	fName := c.FormValue("car_name")
	fImg, err := c.FormFile("image")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  false,
			"message": "bad request",
			"errors":  err.Error(),
			"links": fiber.Map{
				"self": c.OriginalURL(),
			},
		})
	}
	dRate, _ := strconv.ParseFloat(c.FormValue("day_rate"), 32)
	mRate, _ := strconv.ParseFloat(c.FormValue("month_rate"), 32)

	filePath := "./core/presentation/http/resource/img/" + fImg.Filename + time.Now().String()
	c.SaveFile(fImg, filePath)

	if err := h.service.AddCar(fName, float32(dRate), float32(mRate), fImg.Filename); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  false,
			"message": "internal server error",
			"errors":  err.Error(),
			"links": fiber.Map{
				"self": c.OriginalURL(),
			},
		})
	}
	return c.Status(201).JSON(fiber.Map{
		"status":  true,
		"message": "add car success",
		"links": fiber.Map{
			"self": c.OriginalURL(),
		},
	})
}
func (h *CarHandlerImpl) GetCarById(c *fiber.Ctx) error {
	id := c.Params("id")
	car, err := h.service.GetCarByID(id)
	if err != nil {
		if notFoundErr, ok := err.(*pkg.NotFound); ok {
			return c.Status(404).JSON(fiber.Map{
				"status":  false,
				"message": "not found",
				"errors":  notFoundErr.Message,
				"links": fiber.Map{
					"self": c.OriginalURL(),
				},
			})
		} else {
			return c.Status(500).JSON(fiber.Map{
				"status":  false,
				"message": "internal server error",
				"errors":  err.Error(),
				"links": fiber.Map{
					"self": c.OriginalURL(),
				},
			})
		}
	}
	return c.Status(200).JSON(fiber.Map{
		"status":  true,
		"message": "get car success",
		"data":    car,
		"links": fiber.Map{
			"self": c.OriginalURL(),
		},
	})
}
func (h *CarHandlerImpl) UpdateCar(c *fiber.Ctx) error {
	id := c.Params("id")
	body := new(dtos.UpdateRequestCar)
	if err := c.BodyParser(body); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  true,
			"message": "get car success",
			"errors":  err.Error(),
			"links": fiber.Map{
				"self": c.OriginalURL(),
			},
		})
	}
	if err := h.service.UpdateCar(id, *body); err != nil {
		if notFoundErr, ok := err.(*pkg.NotFound); ok {
			return c.Status(404).JSON(fiber.Map{
				"status":  false,
				"message": "not found",
				"errors":  notFoundErr.Message,
				"links": fiber.Map{
					"self": c.OriginalURL(),
				},
			})
		} else {
			return c.Status(500).JSON(fiber.Map{
				"status":  false,
				"message": "internal server error",
				"errors":  err.Error(),
				"links": fiber.Map{
					"self": c.OriginalURL(),
				},
			})
		}
	}
	return c.Status(200).JSON(fiber.Map{
		"status":  true,
		"message": "update car success",
		"links": fiber.Map{
			"self": c.OriginalURL(),
		},
	})
}
func (h *CarHandlerImpl) DeleteCar(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := h.service.DeleteCar(id); err != nil {
		if notFoundErr, ok := err.(*pkg.NotFound); ok {
			return c.Status(404).JSON(fiber.Map{
				"status":  false,
				"message": "not found",
				"errors":  notFoundErr.Message,
				"links": fiber.Map{
					"self": c.OriginalURL(),
				},
			})
		} else {
			return c.Status(500).JSON(fiber.Map{
				"status":  false,
				"message": "internal server error",
				"errors":  err.Error(),
				"links": fiber.Map{
					"self": c.OriginalURL(),
				},
			})
		}
	}
	return c.Status(200).JSON(fiber.Map{
		"status":  true,
		"message": "delete car success",
		"links": fiber.Map{
			"self": c.OriginalURL(),
		},
	})
}
func (h *CarHandlerImpl) GetCars(c *fiber.Ctx) error {
	cars, err := h.service.GetCars()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  false,
			"message": "internal server error",
			"errors":  err.Error(),
			"links": fiber.Map{
				"self": c.OriginalURL(),
			},
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"status":  true,
		"message": "get cars success",
		"data":    cars,
		"links": fiber.Map{
			"self": c.OriginalURL(),
		},
	})
}
