package dto

import (
	"errors"
)

type CreateUpdateActivityReq struct {
	Title *string `json:"title"`
	Email *string `json:"email"`
}

func (r CreateUpdateActivityReq) ValidateCreate() error {
	if r.Title == nil {
		return errors.New("title cannot be null")
	}

	if r.Email == nil {
		return errors.New("email cannot be null")
	}

	return nil
}
