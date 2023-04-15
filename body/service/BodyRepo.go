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

type BodyRepository struct {
	db     *sql.DB
	logger log.Logger
}

func NewBodyRepository(db *sql.DB, logger log.Logger) (IBodyRepository, error) {
	return &BodyRepository{
		db:     db,
		logger: log.With(logger, "repo", "postgresql"),
	}, nil
}

func (e *BodyRepository) GetBodyById(ctx context.Context, id string) (interface{}, error) {
	if id != "bodyaaaaa" {
		return nil, ErrIdNotFound
	}

	body := Body{
		Id:      "bodyaaaa",
		Weight:  55.8,
		Muscle:  20.5,
		FatRate: 32.6,
	}

	return body, nil
}

func (e *BodyRepository) GetBodyByDate(ctx context.Context, date string) (interface{}, error) {
	var bodys []interface{}

	total := Body{
		Id:      "bodyaaaa",
		Weight:  55.8,
		Muscle:  20.5,
		FatRate: 32.6,
	}

	bodys = []interface{}{total}

	return bodys, nil
}

func (e *BodyRepository) GetAllBodys(ctx context.Context) (interface{}, error) {
	var bodys []interface{}

	total := Body{
		Id:      "bodyaaaa",
		Weight:  55.8,
		Muscle:  20.5,
		FatRate: 32.6,
	}

	bodys = []interface{}{total}

	return bodys, nil
}

func (e *BodyRepository) CreateBody(ctx context.Context, body Body) error {
	return nil
}

func (e *BodyRepository) UpdateBody(ctx context.Context, body Body) (string, error) {
	if body.Id != "bodyaaaaa" {
		return "", ErrIdNotFound
	}

	return "Successfully updated", nil
}

func (e *BodyRepository) DeleteBody(ctx context.Context, id string) (string, error) {
	if id != "bodyaaaaa" {
		return "", ErrIdNotFound
	}
	return "Successfully deleted", nil
}
