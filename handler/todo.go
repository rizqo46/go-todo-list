package handler

import (
	"net/http"
	"strconv"

	"github.com/astaxie/beego/orm"
	"github.com/gofiber/fiber/v2"
	"github.com/rizqo46/go-todo-list/dto"
	"github.com/rizqo46/go-todo-list/repo"
)

type TodoHandler struct {
	Orm orm.Ormer
}

func (h TodoHandler) CreateTodo(c *fiber.Ctx) error {
	req := new(dto.CreateUpdateTodoReq)
	if err := c.App().Config().JSONDecoder(c.Body(), req); err != nil {
		return err
	}

	err := req.ValidateCreate()
	if err != nil {
		return BadRequest(c, err.Error())
	}

	todo := &repo.Todo{
		ActivityGroupId: *req.ActivityGroupId,
		Title:           *req.Title,
	}

	const defaultPriority string = "very-high"
	if req.Priority == nil {
		todo.Priority = defaultPriority
	}

	const defaultIsActive bool = true
	if req.IsActive == nil {
		todo.IsActive = defaultIsActive
	}

	err = repo.CreateTodo(h.Orm, todo)
	if err != nil {
		return err
	}

	return Created(c, todo)
}

func (h TodoHandler) GetTodos(c *fiber.Ctx) error {
	activityGroupId := c.Query("activity_group_id")
	activityGroupIdInt := 0
	if activityGroupId != "" {
		activityGroupIdInt, _ = strconv.Atoi(activityGroupId)
	}

	todos, err := repo.GetTodos(h.Orm, activityGroupIdInt)
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}

	return Ok(c, todos)
}

func (h TodoHandler) GetTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	todo, err := repo.GetTodo(h.Orm, id)
	if err == orm.ErrNoRows {
		return NotFound(c, "Todo", id)
	}
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}

	return Ok(c, todo)
}

func (h TodoHandler) UpdateTodo(c *fiber.Ctx) error {
	id := c.Params("id")

	req := new(dto.CreateUpdateTodoReq)
	if err := c.App().Config().JSONDecoder(c.Body(), req); err != nil {
		return err
	}

	todo, err := repo.GetTodo(h.Orm, id)
	if err == orm.ErrNoRows {
		return NotFound(c, "Todo", id)
	}
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}

	updateColumns := make([]string, 0, 4)

	if req.Title != nil {
		todo.Title = *req.Title
		updateColumns = append(updateColumns, "title")
	}

	if req.ActivityGroupId != nil {
		todo.ActivityGroupId = *req.ActivityGroupId
		updateColumns = append(updateColumns, "activity_group_id")
	}

	if req.Priority != nil {
		todo.Priority = *req.Priority
		updateColumns = append(updateColumns, "priority")
	}

	if req.IsActive != nil {
		todo.IsActive = *req.IsActive
		updateColumns = append(updateColumns, "is_active")
	}

	err = repo.UpdateTodo(h.Orm, todo, updateColumns...)
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}

	return Ok(c, todo)
}

func (h TodoHandler) DeleteTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	isExist, err := repo.DeleteTodo(h.Orm, id)
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}

	if !isExist {
		return NotFound(c, "Todo", id)
	}

	return Ok(c, struct{}{})
}
