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
	RepoErr             = errors.New("Unable to handle Repo Request")
	ErrIdNotFound       = errors.New("Id not found")
	ErrPhonenumNotFound = errors.New("Phone num is not found")
)

type FoodRepository struct {
	db     *sql.DB
	logger log.Logger
}

type FoodObj struct {
	// id      uuid.UUID    `db:"id"`
	id      int          `db:"id"`
	name    string       `db:"name"`
	brand   string       `db:"brand"`
	amount  *apd.Decimal `db:"amount"`
	unit    string       `db:"unit"`
	carb    *apd.Decimal `db:"carb"`
	portein *apd.Decimal `db:"portein"`
	fat     *apd.Decimal `db:"fat"`
	cal     *apd.Decimal `db:"cal"`
}

func NewFoodRepository(db *sql.DB, logger log.Logger) (IFoodRepository, error) {
	return &FoodRepository{
		db:     db,
		logger: log.With(logger, "repo", "postgresql"),
	}, nil
}

func newFood(f FoodObj) Food {
	amount, _ := f.amount.Float64()
	carb, _ := f.carb.Float64()
	portein, _ := f.portein.Float64()
	fat, _ := f.fat.Float64()
	cal, _ := f.cal.Float64()

	return Food{
		Foodid:  strconv.Itoa(f.id),
		Name:    f.name,
		Brand:   f.brand,
		Amount:  float32(amount),
		Unit:    f.unit,
		Carb:    float32(carb),
		Portein: float32(portein),
		Fat:     float32(fat),
		Cal:     float32(cal),
	}
}

func (f *FoodRepository) GetFoodById(ctx context.Context, id int64) (interface{}, error) {
	var food FoodObj
	q := `SELECT id, name, brand, amount, unit, carb, portein, fat, cal FROM food WHERE id=$1`
	row := f.db.QueryRowContext(ctx, q, id)
	err := row.Scan(&food.id, &food.name, &food.brand, &food.amount, &food.unit, &food.carb, &food.portein, &food.fat, &food.cal)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrIdNotFound
		}
		return nil, err
	}

	return newFood(food), nil
}

func (f *FoodRepository) GetAllFoods(ctx context.Context) (interface{}, error) {
	var foods []interface{}
	rows, err := f.db.QueryContext(ctx, "SELECT id, name, brand, amount, unit, carb, portein, fat, cal FROM food")
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrIdNotFound
		}
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var food FoodObj
		rows.Scan(&food.id, &food.name, &food.brand, &food.amount, &food.unit, &food.carb, &food.portein, &food.fat, &food.cal)
		foods = append([]interface{}{newFood(food)}, foods...)
	}

	return foods, nil
}

func (f *FoodRepository) CreateFood(ctx context.Context, food Food) error {
	q := `INSERT INTO food (name, brand, amount, unit, carb, portein, fat, cal) values ($1,$2,$3,$4,$5,$6,$7,$8)`
	_, err := f.db.ExecContext(ctx, q, food.Name, food.Brand, food.Amount, food.Unit, food.Cal, food.Portein, food.Fat, food.Cal)
	if err != nil {
		fmt.Println("Error occured inside CreateFood in repo")
		fmt.Println(err.Error())
		return err
	}
	return nil
}

func (f *FoodRepository) UpdateFood(ctx context.Context, food Food) (string, error) {
	fmt.Println(food.Foodid)
	id, _ := strconv.Atoi(food.Foodid)
	fmt.Println(id)
	q := `UPDATE food SET name=$1, brand=$2, amount=$3, unit=$4, carb=$5, portein=$6, fat=$7, cal=$8 WHERE id=$9`
	res, err := f.db.ExecContext(ctx, q, food.Name, food.Brand, food.Amount, food.Unit, food.Cal, food.Portein, food.Fat, food.Cal, id)

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

func (f *FoodRepository) DeleteFood(ctx context.Context, id int) (string, error) {
	res, err := f.db.ExecContext(ctx, "DELETE FROM food WHERE id = $1 ", id)
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
