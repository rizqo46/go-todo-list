package dto

import "errors"

type CreateUpdateTodoReq struct {
	Title           *string `json:"title"`
	ActivityGroupId *int    `json:"activity_group_id"`
	Priority        *string `json:"priority"`
	IsActive        *bool   `json:"is_active"`
}

func (r *CreateUpdateTodoReq) ValidateCreate() error {
	if r.Title == nil {
		return errors.New("title cannot be null")
	}

	if r.ActivityGroupId == nil {
		return errors.New("activity_group_id cannot be null")
	}

	return nil
}
