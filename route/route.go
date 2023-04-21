package route

import (
	"github.com/astaxie/beego/orm"
	"github.com/gofiber/fiber/v2"
	"github.com/rizqo46/go-todo-list/handler"
)

func InitRoute(app *fiber.App, orm orm.Ormer) {
	activityHandler := handler.ActivityHandler{Orm: orm}
	todoHandler := handler.TodoHandler{Orm: orm}

	activityRoutes := app.Group("/activity-groups")
	activityRoutes.Post("", activityHandler.CreateActivity)
	activityRoutes.Get("", activityHandler.GetActivities)
	activityRoutes.Get("/:id", activityHandler.GetActivity)
	activityRoutes.Patch("/:id", activityHandler.UpdateActivity)
	activityRoutes.Delete("/:id", activityHandler.DeleteActivity)

	todoRoutes := app.Group("/todo-items")
	todoRoutes.Get("", todoHandler.GetTodos)
	todoRoutes.Post("", todoHandler.CreateTodo)
	todoRoutes.Get("/:id", todoHandler.GetTodo)
	todoRoutes.Patch("/:id", todoHandler.UpdateTodo)
	todoRoutes.Delete("/:id", todoHandler.DeleteTodo)
}
