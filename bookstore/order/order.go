package order

import (
	"errors"
	"fmt"
)

type order struct {
	id string
}

func New(id string) (*order, error) {
	if id == "" {
		return &order{}, fmt.Errorf("invalid ID: %q", id)
	}
	return &order{
		id: id,
	}, nil
}

func (o order) ID() string {
	return o.id
}

func (o *order) SetID(id string) error {
	if id == "" {
		return errors.New("Invalid ID! The ID must not be empty")
	}
	o.id = id
	return nil
}
