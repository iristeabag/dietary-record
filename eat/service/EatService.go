package service

import (
	"context"

	"github.com/go-kit/log"
)

type Eat struct {
	Id      string  `json:"eatid"`
	Foodid  string  `json:"foodid"`
	Name    string  `json:"name"`
	Amount  float64 `json:"amount"`
	Unit    string  `json:"unit"`
	Carb    float64 `json:"carb"`
	Portein float64 `json:"portein"`
	Fat     float64 `json:"fat"`
	Cal     float64 `json:"cal"`
}

type TotalEat struct {
	Carb    float64 `json:"carb"`
	Portein float64 `json:"portein"`
	Fat     float64 `json:"fat"`
	Cal     float64 `json:"cal"`
}

type IEatRepository interface {
	GetEatById(ctx context.Context, id string) (interface{}, error)
	GetTotalEatByDate(ctx context.Context, date string) (interface{}, error)
	GetEatsByDate(ctx context.Context, date string) (interface{}, error)
	GetAllEats(ctx context.Context) (interface{}, error)
	CreateEat(ctx context.Context, eat Eat) error
	UpdateEat(ctx context.Context, eat Eat) (string, error)
	DeleteEat(ctx context.Context, id string) (string, error)
}

type EatService struct {
	repository IEatRepository
	logger     log.Logger
}

type IEatService interface {
	GetEatById(ctx context.Context, id string) (interface{}, error)
	GetTotalEatByDate(ctx context.Context, date string) (interface{}, error)
	GetEatsByDate(ctx context.Context, date string) (interface{}, error)
	GetAllEats(ctx context.Context) (interface{}, error)
	CreateEat(ctx context.Context, eat Eat) (string, error)
	UpdateEat(ctx context.Context, eat Eat) (string, error)
	DeleteEat(ctx context.Context, id string) (string, error)
}

func NewEatService(rep IEatRepository, logger log.Logger) IEatService {
	return &EatService{
		repository: rep,
		logger:     logger,
	}
}

func (s EatService) GetEatById(ctx context.Context, id string) (interface{}, error) {
	var eat interface{}
	var empty interface{}

	eat, err := s.repository.GetEatById(ctx, id)
	if err == nil {
		return eat, nil
	}

	return empty, err
}

func (s EatService) GetTotalEatByDate(ctx context.Context, date string) (interface{}, error) {
	var eats interface{}
	var empty interface{}

	eats, err := s.repository.GetTotalEatByDate(ctx, date)
	if err == nil {
		return eats, nil
	}

	return empty, err
}

func (s EatService) GetEatsByDate(ctx context.Context, date string) (interface{}, error) {
	var eats interface{}
	var empty interface{}

	eats, err := s.repository.GetEatsByDate(ctx, date)
	if err == nil {
		return eats, nil
	}

	return empty, err
}

func (s EatService) GetAllEats(ctx context.Context) (interface{}, error) {
	var eats interface{}
	var empty interface{}

	eats, err := s.repository.GetAllEats(ctx)
	if err == nil {
		return eats, nil
	}

	return empty, err
}

func (s EatService) CreateEat(ctx context.Context, eat Eat) (string, error) {
	var msg = "success"
	var err error

	eatDetail := Eat{
		Id:      eat.Id,
		Foodid:  eat.Foodid,
		Name:    eat.Name,
		Amount:  eat.Amount,
		Unit:    eat.Unit,
		Carb:    eat.Carb,
		Portein: eat.Portein,
		Fat:     eat.Fat,
		Cal:     eat.Cal,
	}

	if err := s.repository.CreateEat(ctx, eatDetail); err != nil {
		return "", err
	}

	return msg, err
}

func (s EatService) UpdateEat(ctx context.Context, eat Eat) (string, error) {
	eatDetail := Eat{
		Id:      eat.Id,
		Foodid:  eat.Foodid,
		Name:    eat.Name,
		Amount:  eat.Amount,
		Unit:    eat.Unit,
		Carb:    eat.Carb,
		Portein: eat.Portein,
		Fat:     eat.Fat,
		Cal:     eat.Cal,
	}

	msg, err := s.repository.UpdateEat(ctx, eatDetail)
	if err != nil {
		return "", err
	}

	return msg, nil
}

func (s EatService) DeleteEat(ctx context.Context, id string) (string, error) {
	msg, err := s.repository.DeleteEat(ctx, id)
	if err != nil {
		return "", err
	}

	return msg, nil
}