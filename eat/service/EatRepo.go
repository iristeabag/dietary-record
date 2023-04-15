package service

import (
	"context"
	"database/sql"
	"errors"

	"github.com/go-kit/log"
)

var (
	ErrIdNotFound = errors.New("id not found")
)

type EatRepository struct {
	db     *sql.DB
	logger log.Logger
}

func NewEatRepository(db *sql.DB, logger log.Logger) (IEatRepository, error) {
	return &EatRepository{
		db:     db,
		logger: log.With(logger, "repo", "postgresql"),
	}, nil
}

func (e *EatRepository) GetEatById(ctx context.Context, id string) (interface{}, error) {
	if id != "eataaaaa" {
		return nil, ErrIdNotFound
	}

	eat := Eat{
		Id:      "eataaaaa",
		Foodid:  "foodid",
		Name:    "三明治",
		Amount:  1,
		Unit:    "份",
		Carb:    20.6,
		Portein: 15,
		Fat:     9.5,
		Cal:     230,
	}

	return eat, nil
}

func (e *EatRepository) GetTotalEatByDate(ctx context.Context, date string) (interface{}, error) {

	total := TotalEat{
		Carb:    20.6,
		Portein: 15,
		Fat:     9.5,
		Cal:     230,
	}

	return total, nil
}

func (e *EatRepository) GetEatsByDate(ctx context.Context, date string) (interface{}, error) {
	var eats []interface{}

	total := TotalEat{
		Carb:    12,
		Portein: 15,
		Fat:     9.5,
		Cal:     230,
	}

	eats = []interface{}{total}

	return eats, nil
}

func (e *EatRepository) GetAllEats(ctx context.Context) (interface{}, error) {
	var eats []interface{}

	total := TotalEat{
		Carb:    10,
		Portein: 15,
		Fat:     9.5,
		Cal:     230,
	}

	eats = []interface{}{total}

	return eats, nil
}

func (e *EatRepository) CreateEat(ctx context.Context, eat Eat) error {
	return nil
}

func (e *EatRepository) UpdateEat(ctx context.Context, eat Eat) (string, error) {
	if eat.Id != "eataaaaa" {
		return "", ErrIdNotFound
	}

	return "Successfully updated", nil
}

func (e *EatRepository) DeleteEat(ctx context.Context, id string) (string, error) {
	if id != "eataaaaa" {
		return "", ErrIdNotFound
	}
	return "Successfully deleted", nil
}
