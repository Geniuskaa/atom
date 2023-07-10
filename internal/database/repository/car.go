package repository

import (
	"atom/internal/model"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

type carRepo struct {
	pool *pgxpool.Pool
}

func NewCarRepo(pool *pgxpool.Pool) *carRepo {
	return &carRepo{pool: pool}
}

type Car interface {
	RegisterCar(ctx context.Context, car *model.Car) (int, error)
	UpdateCarInfo(ctx context.Context, car *model.Car, id int) error
	DeleteCarInfo(ctx context.Context, id int) error
	GetCarInfo(ctx context.Context, plate string) (*model.Car, error)
	GetAllCarsInfo(ctx context.Context) ([]model.Car, error)
	GetCarMileage(ctx context.Context, plate string) (int, error)
	AddCarMileage(ctx context.Context, id int, km int) error
}

func (c *carRepo) RegisterCar(ctx context.Context, car *model.Car) (int, error) {
	t, err := time.Parse("2006", car.YearOfManufacture)
	if err != nil {
		return -1, fmt.Errorf("RegisterCar failed: %w", err)
	}

	timeStamp := pgtype.Timestamp{}
	err = timeStamp.Scan(t)
	if err != nil {
		return -1, fmt.Errorf("RegisterCar failed: %w", err)
	}

	row := c.pool.QueryRow(ctx, `INSERT INTO car_info (plate, model, year_of_manufacture, type) VALUES ($1, $2, $3, $4) 
        RETURNING id`, car.Plate, car.Model, timeStamp, car.Type.String())

	id := -1
	err = row.Scan(&id)
	if err != nil {
		return -1, fmt.Errorf("RegisterCar failed: %w", err)
	}

	return id, nil
}

func (c *carRepo) UpdateCarInfo(ctx context.Context, car *model.Car, id int) error {
	t, err := time.Parse("2006", car.YearOfManufacture)
	if err != nil {
		return fmt.Errorf("UpdateCarInfo failed: %w", err)
	}

	timeStamp := pgtype.Timestamp{}
	err = timeStamp.Scan(t)
	if err != nil {
		return fmt.Errorf("UpdateCarInfo failed: %w", err)
	}

	_, err = c.pool.Exec(ctx, `UPDATE car_info SET plate=$1, model=$2, year_of_manufacture=$3, type=$4 WHERE id=$5;`,
		car.Plate, car.Model, timeStamp, car.Type, id)
	if err != nil {
		return fmt.Errorf("UpdateCarInfo failed: %w", err)
	}

	return nil
}

func (c *carRepo) GetCarInfo(ctx context.Context, plate string) (*model.Car, error) {
	row := c.pool.QueryRow(ctx, `SELECT plate, model, year_of_manufacture, type, is_deleted FROM car_info WHERE plate=$1`, plate)

	car := model.Car{}
	t := pgtype.Timestamp{}
	carType := ""
	isDeleted := false
	err := row.Scan(&car.Plate, &car.Model, &t, &carType, &isDeleted)
	if err != nil {
		return nil, fmt.Errorf("GetCarInfo failed: %w", err)
	}

	if isDeleted {
		return nil, fmt.Errorf("GetCarInfo failed: %w", model.ErrCarInfoDeleted)
	}

	car.SetCarType(carType)
	car.SetYear(t.Time)

	return &car, nil
}

func (c *carRepo) DeleteCarInfo(ctx context.Context, id int) error {
	_, err := c.pool.Exec(ctx, `UPDATE car_info SET is_deleted=true WHERE id=$1;`, id)
	if err != nil {
		return fmt.Errorf("DeleteCarInfo failed: %w", err)
	}

	return nil
}

func (c *carRepo) GetAllCarsInfo(ctx context.Context) ([]model.Car, error) {
	rows, err := c.pool.Query(ctx, `SELECT plate, model, year_of_manufacture, type FROM car_info WHERE is_deleted=false;`)
	if err != nil {
		return nil, fmt.Errorf("GetAllCarsInfo failed: %w", err)
	}

	arr := make([]model.Car, 0, 20)
	for rows.Next() {
		car := model.Car{}
		t := pgtype.Timestamp{}
		carType := ""

		err := rows.Scan(&car.Plate, &car.Model, &t, &carType)
		if err != nil {
			continue
		}

		car.SetYear(t.Time)
		car.SetCarType(carType)
		arr = append(arr, car)
	}

	return arr, nil
}

func (c *carRepo) GetCarMileage(ctx context.Context, plate string) (int, error) {
	row := c.pool.QueryRow(ctx, `SELECT mileage FROM car_mileage WHERE plate=$1;`, plate)

	mileage := -1
	err := row.Scan(&mileage)
	if err != nil {
		return mileage, fmt.Errorf("GetCarMileage failed: %w", err)
	}

	return mileage, nil
}

func (c *carRepo) AddCarMileage(ctx context.Context, id int, km int) error {
	row := c.pool.QueryRow(ctx, `WITH car_plate AS (SELECT plate FROM car_info WHERE id=$1) SELECT count(*) 
		FROM car_mileage WHERE plate = (SELECT plate FROM car_plate);`, id)
	count := 0

	err := row.Scan(&count)
	if err != nil {
		return fmt.Errorf("AddCarMileage failed: %w", err)
	}

	if count == 0 {
		_, err := c.pool.Exec(ctx, `WITH car_plate AS (SELECT plate FROM car_info WHERE id=$1) INSERT INTO car_mileage 
			(plate, mileage) SELECT plate, $2 FROM car_plate;`, id, km)
		if err != nil {
			return fmt.Errorf("AddCarMileage failed: %w", err)
		}
	}

	if count > 0 {
		_, err := c.pool.Exec(ctx, `WITH car_plate AS (SELECT plate FROM car_info WHERE id=$1) UPDATE car_mileage 
		SET mileage=car_mileage.mileage + $2 WHERE plate = (SELECT plate FROM car_plate);`, id, km)
		if err != nil {
			return fmt.Errorf("AddCarMileage failed: %w", err)
		}
	}

	return nil
}
