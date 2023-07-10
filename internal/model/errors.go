package model

import "errors"

var (
	ErrCarType        = errors.New("Car type does not match the allowed values")
	ErrCarInfoDeleted = errors.New("Client can`t get car info, because it was deleted")
)
