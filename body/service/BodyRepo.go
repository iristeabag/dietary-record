package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strconv"

	"github.com/cockroachdb/apd"
	"github.com/go-kit/log"
)

var (
	ErrIdNotFound = errors.New("id not found")
)

type BodyRepository struct {
	db     *sql.DB
	logger log.Logger
}

type BodyObj struct {
	// id      uuid.UUID    `db:"id"`
	id       int          `db:"id"`
	weight   *apd.Decimal `db:"amount"`
	muscle   *apd.Decimal `db:"carb"`
	fat_rate *apd.Decimal `db:"portein"`
}

func NewBodyRepository(db *sql.DB, logger log.Logger) (IBodyRepository, error) {
	return &BodyRepository{
		db:     db,
		logger: log.With(logger, "repo", "postgresql"),
	}, nil
}

func newBody(e BodyObj) Body {
	weight, _ := e.weight.Float64()
	muscle, _ := e.muscle.Float64()
	fat_rate, _ := e.fat_rate.Float64()

	return Body{
		Id:      strconv.Itoa(e.id),
		Weight:  weight,
		Muscle:  muscle,
		FatRate: fat_rate,
	}
}

func (e *BodyRepository) GetBodyById(ctx context.Context, id int) (interface{}, error) {
	var body BodyObj
	q := `SELECT id, weight, muscle, fat_rate FROM body WHERE id=$1`
	row := e.db.QueryRowContext(ctx, q, id)
	err := row.Scan(&body.id, &body.weight, &body.muscle, &body.fat_rate)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrIdNotFound
		}
		return nil, err
	}

	return newBody(body), nil
}

func (e *BodyRepository) GetAllBodys(ctx context.Context) (interface{}, error) {
	var bodys []interface{}
	rows, err := e.db.QueryContext(ctx, "SELECT id, weight, muscle, fat_rate FROM body")
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrIdNotFound
		}
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var body BodyObj
		rows.Scan(&body.id, &body.weight, &body.muscle, &body.fat_rate)
		bodys = append([]interface{}{newBody(body)}, bodys...)
	}

	return bodys, nil
}

func (e *BodyRepository) CreateBody(ctx context.Context, body Body) error {
	q := `INSERT INTO body (weight, muscle, fat_rate) values ($1,$2,$3)`
	_, err := e.db.ExecContext(ctx, q, body.Weight, body.Muscle, body.FatRate)
	if err != nil {
		fmt.Println("Error occured inside CreateBody in repo")
		fmt.Println(err.Error())
		return err
	}
	return nil
}

func (e *BodyRepository) UpdateBody(ctx context.Context, body Body) (string, error) {
	id, _ := strconv.Atoi(body.Id)
	fmt.Println(id)
	q := `UPDATE body SET weight=$1, muscle=$2, fat_rate=$3 WHERE id=$4`
	res, err := e.db.ExecContext(ctx, q, body.Weight, body.Muscle, body.FatRate, body.Id)

	if err != nil {
		fmt.Println("Error occured inside UpdateFood in repo")
		fmt.Println(err.Error())
		return "", err
	}

	rowCnt, err := res.RowsAffected()
	if err != nil {
		return "", err
	}
	if rowCnt == 0 {
		return "", ErrIdNotFound
	}

	return "Successfully updated", nil
}

func (e *BodyRepository) DeleteBody(ctx context.Context, id int) (string, error) {
	res, err := e.db.ExecContext(ctx, "DELETE FROM body WHERE id = $1 ", id)
	if err != nil {
		return "", err
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		return "", err
	} else if rowCnt == 0 {
		return "", ErrIdNotFound
	}

	return "Successfully deleted", nil
}
