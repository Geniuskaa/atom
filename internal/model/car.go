package model

import (
	"strconv"
	"strings"
	"time"
)

type carType int64

const (
	Undefined carType = iota
	Carsharing
	Taxi
	Delivery
)

func newCarType(t string) carType {
	t = strings.ToUpper(t)

	switch t {
	case Carsharing.String():
		return Carsharing
	case Taxi.String():
		return Taxi
	case Delivery.String():
		return Delivery
	default:
		return Undefined
	}
}

func (c carType) String() string {
	switch c {
	case Carsharing:
		return strings.ToUpper("carsharing")
	case Taxi:
		return strings.ToUpper("taxi")
	case Delivery:
		return strings.ToUpper("delivery")
	default:
		return strings.ToUpper("undefined")
	}
}

func (c carType) IsValid() bool {
	switch c {
	case Carsharing, Taxi, Delivery:
		return true
	}
	return false
}

type Car struct {
	Plate             string
	Model             string
	YearOfManufacture string
	Type              carType
}

// Заменяет такую конструкцию: car.YearOfManufacture = strconv.Itoa(t.Time.Year())
func (c *Car) SetYear(year time.Time) {
	c.YearOfManufacture = strconv.Itoa(year.Year())
}

func (c *Car) SetCarType(t string) {
	c.Type = newCarType(t)
}
