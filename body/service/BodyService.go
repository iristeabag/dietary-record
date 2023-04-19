package service

import (
	"context"

	"github.com/go-kit/log"
)

type Body struct {
	Id      string  `json:"bodyid"`
	Weight  float64 `json:"weight"`
	Muscle  float64 `json:"muscle"`
	FatRate float64 `json:"fatrate"`
}

type IBodyRepository interface {
	GetBodyById(ctx context.Context, id int) (interface{}, error)
	GetAllBodys(ctx context.Context) (interface{}, error)
	CreateBody(ctx context.Context, body Body) error
	UpdateBody(ctx context.Context, body Body) (string, error)
	DeleteBody(ctx context.Context, id int) (string, error)
}

type BodyService struct {
	repository IBodyRepository
	logger     log.Logger
}

type IBodyService interface {
	GetBodyById(ctx context.Context, id int) (interface{}, error)
	GetAllBodys(ctx context.Context) (interface{}, error)
	CreateBody(ctx context.Context, body Body) (string, error)
	UpdateBody(ctx context.Context, body Body) (string, error)
	DeleteBody(ctx context.Context, id int) (string, error)
}

func NewBodyService(rep IBodyRepository, logger log.Logger) IBodyService {
	return &BodyService{
		repository: rep,
		logger:     logger,
	}
}

func (s BodyService) GetBodyById(ctx context.Context, id int) (interface{}, error) {
	var body interface{}
	var empty interface{}

	body, err := s.repository.GetBodyById(ctx, id)
	if err == nil {
		return body, nil
	}

	return empty, err
}

func (s BodyService) GetAllBodys(ctx context.Context) (interface{}, error) {
	var bodys interface{}
	var empty interface{}

	bodys, err := s.repository.GetAllBodys(ctx)
	if err == nil {
		return bodys, nil
	}

	return empty, err
}

func (s BodyService) CreateBody(ctx context.Context, body Body) (string, error) {
	var msg = "success"
	var err error

	bodyDetail := Body{
		Id:      body.Id,
		Weight:  body.Weight,
		Muscle:  body.Muscle,
		FatRate: body.FatRate,
	}

	if err := s.repository.CreateBody(ctx, bodyDetail); err != nil {
		return "", err
	}

	return msg, err
}

func (s BodyService) UpdateBody(ctx context.Context, body Body) (string, error) {

	msg, err := s.repository.UpdateBody(ctx, body)
	if err != nil {
		return "", err
	}

	return msg, nil
}

func (s BodyService) DeleteBody(ctx context.Context, id int) (string, error) {
	msg, err := s.repository.DeleteBody(ctx, id)
	if err != nil {
		return "", err
	}

	return msg, nil
}
