package service

import (
	"context"
	"errors"

	"github.com/go-kit/log"
)

var (
	errInvalidFoodID = errors.New("food id is invalid")
)

type Food struct {
	Foodid  string  `json:"foodid"`
	Name    string  `json:"name"`
	Brand   string  `json:"brand"`
	Amount  float64 `json:"amount"`
	Unit    string  `json:"unit"`
	Carb    float64 `json:"carb"`
	Portein float64 `json:"portein"`
	Fat     float64 `json:"fat"`
	Cal     float64 `json:"cal"`
}

type IFoodRepository interface {
	GetFoodById(ctx context.Context, id int) (interface{}, error)
	GetAllFoods(ctx context.Context) (interface{}, error)
	CreateFood(ctx context.Context, food Food) error
	UpdateFood(ctx context.Context, food Food) (string, error)
	DeleteFood(ctx context.Context, id int) (string, error)
}

type FoodService struct {
	repository IFoodRepository
	logger     log.Logger
}

type IFoodService interface {
	GetFoodById(ctx context.Context, id int) (interface{}, error)
	GetAllFoods(ctx context.Context) (interface{}, error)
	CreateFood(ctx context.Context, food Food) (string, error)
	UpdateFood(ctx context.Context, food Food) (string, error)
	DeleteFood(ctx context.Context, id int) (string, error)
}

// NewService creates and returns a new Account service instance
func NewFoodService(rep IFoodRepository, logger log.Logger) IFoodService {
	return &FoodService{
		repository: rep,
		logger:     logger,
	}
}

func (s FoodService) GetFoodById(ctx context.Context, id int) (interface{}, error) {
	var food interface{}
	var empty interface{}

	food, err1 := s.repository.GetFoodById(ctx, id)
	if err1 != nil {
		return empty, err1
	}

	return food, nil
}

func (s FoodService) GetAllFoods(ctx context.Context) (interface{}, error) {
	var foods interface{}
	var empty interface{}

	foods, err := s.repository.GetAllFoods(ctx)
	if err == nil {
		return foods, nil
	}

	return empty, err
}

func (s FoodService) CreateFood(ctx context.Context, food Food) (string, error) {
	var msg = "success"

	foodDetail := Food{
		Name:    food.Name,
		Brand:   food.Brand,
		Amount:  food.Amount,
		Unit:    food.Unit,
		Carb:    food.Carb,
		Portein: food.Portein,
		Fat:     food.Fat,
		Cal:     food.Cal,
	}

	if err := s.repository.CreateFood(ctx, foodDetail); err != nil {
		return "", err
	}

	return msg, nil
}

func (s FoodService) UpdateFood(ctx context.Context, food Food) (string, error) {
	// foodDetail := Food{
	// 	Foodid:  id,
	// 	Name:    food.Name,
	// 	Brand:   food.Brand,
	// 	Amount:  food.Amount,
	// 	Unit:    food.Unit,
	// 	Carb:    food.Carb,
	// 	Portein: food.Portein,
	// 	Fat:     food.Fat,
	// 	Cal:     food.Cal,
	// }

	msg, err := s.repository.UpdateFood(ctx, food)
	if err != nil {
		return "", err
	}

	return msg, nil
}

func (s FoodService) DeleteFood(ctx context.Context, id int) (string, error) {
	msg, err := s.repository.DeleteFood(ctx, id)
	if err != nil {
		return "", err
	}

	return msg, nil
}
