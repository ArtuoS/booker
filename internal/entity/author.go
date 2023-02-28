package entity

import (
	"errors"

	"github.com/google/uuid"
)

type Author struct {
	Id   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

func NewAuthor(name string) Author {
	return Author{
		Id:   uuid.New(),
		Name: name,
	}
}

func (a *Author) Validate() error {
	if a.Name == "" {
		return errors.New("author name is empty")
	}

	return nil
}
