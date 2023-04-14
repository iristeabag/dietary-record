package service

import (
	"context"
	"database/sql"
	"errors"

	"github.com/go-kit/log"
)

var (
	RepoErr             = errors.New("Unable to handle Repo Request")
	ErrIdNotFound       = errors.New("Id not found")
	ErrPhonenumNotFound = errors.New("Phone num is not found")
)

type FoodRepository struct {
	db     *sql.DB
	logger log.Logger
}

func NewFoodRepository(db *sql.DB, logger log.Logger) (IFoodRepository, error) {
	return &FoodRepository{
		db:     db,
		logger: log.With(logger, "repo", "postgresql"),
	}, nil
}

func (f *FoodRepository) GetFoodById(ctx context.Context, id string) (interface{}, error) {
	if id != "aaaaa" {
		return nil, ErrIdNotFound
	}

	food := Food{
		Foodid:  "aaaaa",
		Name:    "三明治",
		Brand:   "7-11",
		Amount:  1,
		Unit:    "份",
		Carb:    20.6,
		Portein: 15,
		Fat:     9.5,
		Cal:     230,
	}

	return food, nil
}

func (f *FoodRepository) GetAllFoods(ctx context.Context) (interface{}, error) {
	var foods []interface{}

	return foods, nil
}

func (f *FoodRepository) CreateFood(ctx context.Context, food Food) error {
	return nil
}

func (f *FoodRepository) UpdateFood(ctx context.Context, food Food) (string, error) {
	if food.Foodid != "aaaaa" {
		return "", ErrIdNotFound
	}

	return "Successfully updated", nil
}

func (f *FoodRepository) DeleteFood(ctx context.Context, id string) (string, error) {
	if id != "aaaaa" {
		return "", ErrIdNotFound
	}
	return "Successfully deleted", nil
}
