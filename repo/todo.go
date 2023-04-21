package repo

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type Todo struct {
	TodoId          int       `orm:"auto" json:"id"`
	ActivityGroupId int       `json:"activity_group_id"`
	Title           string    `json:"title"`
	Priority        string    `json:"priority"`
	IsActive        bool      `json:"is_active"`
	CreatedAt       time.Time `json:"created_at" orm:"auto_now_add;type(datetime)"`
	UpdatedAt       time.Time `json:"updated_at" orm:"auto_now;type(datetime)"`
}

const todosTableName string = "todos"

func (u *Todo) TableName() string {
	return todosTableName
}

func CreateTodo(o orm.Ormer, todo *Todo) error {
	_, err := o.Insert(todo)
	return err
}

func GetTodos(o orm.Ormer, activityGroupId int) ([]Todo, error) {
	todos := []Todo{}
	qs := o.QueryTable(todosTableName)

	if activityGroupId > 0 {
		qs = qs.Filter("activity_group_id", activityGroupId)
	}

	_, err := qs.All(&todos)
	if err != nil {
		return nil, err
	}

	return todos, nil
}

func GetTodo(o orm.Ormer, id string) (*Todo, error) {
	todo := Todo{}
	err := o.QueryTable(todosTableName).
		Filter("todo_id", id).
		One(&todo)
	if err != nil {
		return nil, err
	}

	return &todo, nil
}

func UpdateTodo(o orm.Ormer, todo *Todo, columns ...string) error {
	_, err := o.Update(todo, columns...)
	return err
}

func DeleteTodo(o orm.Ormer, id string) (isExist bool, err error) {
	idExist, err := o.QueryTable(todosTableName).
		Filter("todo_id", id).
		Delete()
	if err != nil {
		return false, err
	}

	return idExist != 0, nil
}
