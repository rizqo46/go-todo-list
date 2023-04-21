package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func Ok(c *fiber.Ctx, data interface{}) error {
	return c.JSON(Response{
		Status:  "Success",
		Message: "Success",
		Data:    data,
	})
}

func Created(c *fiber.Ctx, data interface{}) error {
	return c.
		Status(http.StatusCreated).
		JSON(Response{
			Status:  "Success",
			Message: "Success",
			Data:    data,
		})
}

func NotFound(c *fiber.Ctx, dataName string, id string) error {
	return c.
		Status(http.StatusNotFound).
		JSON(Response{
			Status:  "Not Found",
			Message: dataName + " with ID " + id + " Not Found",
			Data:    struct{}{},
		})
}

func BadRequest(c *fiber.Ctx, message string) error {
	return c.
		Status(http.StatusBadRequest).
		JSON(Response{
			Status:  "Bad Request",
			Message: message,
			Data:    struct{}{},
		})
}

func InvalidRequest(c *fiber.Ctx) error {
	return c.
		Status(http.StatusBadRequest).
		JSON(Response{
			Status:  "Invalid Request",
			Message: "Invalid Request",
			Data:    struct{}{},
		})
}
