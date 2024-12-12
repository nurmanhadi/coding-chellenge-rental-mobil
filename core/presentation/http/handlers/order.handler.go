package handlers

import (
	"coding-chelleng/core/application/dtos"
	"coding-chelleng/core/application/services"
	"coding-chelleng/pkg"

	"github.com/gofiber/fiber/v2"
)

type OrderHandler interface {
	AddOrder(c *fiber.Ctx) error
	DeleteOrder(c *fiber.Ctx) error
	GetOrders(c *fiber.Ctx) error
	UpdateOrders(c *fiber.Ctx) error
}

type OrderHandlerImpl struct {
	service services.OrderService
}

func NewOrderHandler(service *services.OrderService) OrderHandler {
	return &OrderHandlerImpl{service: *service}
}
func (h *OrderHandlerImpl) AddOrder(c *fiber.Ctx) error {
	body := new(dtos.AddRequestOrder)
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
	if err := h.service.AddOrder(*body); err != nil {
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
	return c.Status(201).JSON(fiber.Map{
		"status":  true,
		"message": "add oerder success",
		"links": fiber.Map{
			"self": c.OriginalURL(),
		},
	})
}
func (h *OrderHandlerImpl) UpdateOrders(c *fiber.Ctx) error {
	id := c.Params("id")
	body := new(dtos.UpdateRequestOrder)
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
	if err := h.service.UpdateOrder(id, *body); err != nil {
		if notFoundErr, ok := err.(*pkg.NotFound); ok {
			return c.Status(404).JSON(fiber.Map{
				"status":  false,
				"message": "not found",
				"errors":  notFoundErr.Message,
				"links": fiber.Map{
					"self": c.OriginalURL(),
				},
			})
		} else if bd, ok := err.(*pkg.BadRequest); ok {
			return c.Status(400).JSON(fiber.Map{
				"status":  false,
				"message": "not found",
				"errors":  bd.Message,
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
		"message": "update order success",
		"links": fiber.Map{
			"self": c.OriginalURL(),
		},
	})
}

func (h *OrderHandlerImpl) DeleteOrder(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := h.service.DeleteOrder(id); err != nil {
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
		"message": "delete oerder success",
		"links": fiber.Map{
			"self": c.OriginalURL(),
		},
	})
}
func (h *OrderHandlerImpl) GetOrders(c *fiber.Ctx) error {
	orders, err := h.service.GetOrders()
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
		"message": "get orders success",
		"data":    orders,
		"links": fiber.Map{
			"self": c.OriginalURL(),
		},
	})
}
