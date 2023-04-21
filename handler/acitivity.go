package handler

import (
	"net/http"

	"github.com/astaxie/beego/orm"
	"github.com/gofiber/fiber/v2"
	"github.com/rizqo46/go-todo-list/dto"
	"github.com/rizqo46/go-todo-list/repo"
)

type ActivityHandler struct {
	Orm orm.Ormer
}

func (h ActivityHandler) CreateActivity(c *fiber.Ctx) error {
	req := new(dto.CreateUpdateActivityReq)

	if err := c.App().Config().JSONDecoder(c.Body(), req); err != nil {
		return err
	}

	err := req.ValidateCreate()
	if err != nil {
		return BadRequest(c, err.Error())
	}

	activity := &repo.Activity{
		Title: *req.Title,
		Email: *req.Email,
	}

	err = repo.CreateActivity(h.Orm, activity)
	if err != nil {
		return err
	}

	return Created(c, activity)
}

func (h ActivityHandler) GetActivities(c *fiber.Ctx) error {
	activity, err := repo.GetActivities(h.Orm)
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}

	return Ok(c, activity)
}

func (h ActivityHandler) GetActivity(c *fiber.Ctx) error {
	id := c.Params("id")

	activity, err := repo.GetActivity(h.Orm, id)
	if err == orm.ErrNoRows {
		return NotFound(c, "Activity", id)
	}

	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}

	return Ok(c, activity)
}

func (h ActivityHandler) UpdateActivity(c *fiber.Ctx) error {
	id := c.Params("id")

	req := new(dto.CreateUpdateActivityReq)
	if err := c.App().Config().JSONDecoder(c.Body(), req); err != nil {
		return err
	}

	activity, err := repo.GetActivity(h.Orm, id)
	if err == orm.ErrNoRows {
		return NotFound(c, "Activity", id)
	}
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}

	updateColumns := make([]string, 0, 2)

	if req.Email != nil {
		activity.Email = *req.Email
		updateColumns = append(updateColumns, "email")
	}

	if req.Title != nil {
		activity.Title = *req.Title
		updateColumns = append(updateColumns, "title")
	}

	err = repo.UpdateActivity(h.Orm, activity, updateColumns...)
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}

	return Ok(c, activity)
}

func (h ActivityHandler) DeleteActivity(c *fiber.Ctx) error {
	id := c.Params("id")
	isExist, err := repo.DeleteActivity(h.Orm, id)
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}

	if !isExist {
		return NotFound(c, "Activity", id)
	}

	return Ok(c, struct{}{})
}
