package creditcard

import (
	"errors"
	"fmt"
)

type card struct {
	number string
}

func New(num string) (*card, error) {
	if num == "" {
		return &card{}, fmt.Errorf("invalid N: \"%q\"", num)
	}
	return &card{
		number: num,
	}, nil
}

func (c card) Number() string {
	return c.number
}

func (c *card) SetNumber(num string) error {
	if num == "" {
		return errors.New("invalid number (must not be empty)")
	}
	c.number = num
	return nil
}
