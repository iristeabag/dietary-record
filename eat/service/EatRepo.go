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

type EatRepository struct {
	db     *sql.DB
	logger log.Logger
}

type EatObj struct {
	// id      uuid.UUID    `db:"id"`
	id      int          `db:"id"`
	food_id int          `db:"food_id"`
	name    string       `db:"name"`
	amount  *apd.Decimal `db:"amount"`
	unit    string       `db:"unit"`
	carb    *apd.Decimal `db:"carb"`
	portein *apd.Decimal `db:"portein"`
	fat     *apd.Decimal `db:"fat"`
	cal     *apd.Decimal `db:"cal"`
}

func NewEatRepository(db *sql.DB, logger log.Logger) (IEatRepository, error) {
	return &EatRepository{
		db:     db,
		logger: log.With(logger, "repo", "postgresql"),
	}, nil
}

func newEat(e EatObj) Eat {
	amount, _ := e.amount.Float64()
	carb, _ := e.carb.Float64()
	portein, _ := e.portein.Float64()
	fat, _ := e.fat.Float64()
	cal, _ := e.cal.Float64()

	return Eat{
		Id:      strconv.Itoa(e.id),
		Foodid:  strconv.Itoa(e.food_id),
		Name:    e.name,
		Amount:  amount,
		Unit:    e.unit,
		Carb:    carb,
		Portein: portein,
		Fat:     fat,
		Cal:     cal,
	}
}

func (e *EatRepository) GetEatById(ctx context.Context, id int) (interface{}, error) {
	var eat EatObj
	q := `SELECT id, food_id, name, amount, unit, carb, portein, fat, cal FROM eat WHERE id=$1`
	row := e.db.QueryRowContext(ctx, q, id)
	err := row.Scan(&eat.id, &eat.food_id, &eat.name, &eat.amount, &eat.unit, &eat.carb, &eat.portein, &eat.fat, &eat.cal)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrIdNotFound
		}
		return nil, err
	}

	return newEat(eat), nil
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
	rows, err := e.db.QueryContext(ctx, "SELECT id, food_id, name, amount, unit, carb, portein, fat, cal FROM eat")
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrIdNotFound
		}
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var eat EatObj
		rows.Scan(&eat.id, &eat.food_id, &eat.name, &eat.amount, &eat.unit, &eat.carb, &eat.portein, &eat.fat, &eat.cal)
		eats = append([]interface{}{newEat(eat)}, eats...)
	}

	return eats, nil
}

func (e *EatRepository) CreateEat(ctx context.Context, eat Eat) error {
	fmt.Println("comming")
	Foodid, _ := strconv.Atoi(eat.Foodid)
	q := `INSERT INTO eat (food_id, name, amount, unit, carb, portein, fat, cal) values ($1,$2,$3,$4,$5,$6,$7,$8)`
	_, err := e.db.ExecContext(ctx, q, Foodid, eat.Name, eat.Amount, eat.Unit, eat.Cal, eat.Portein, eat.Fat, eat.Cal)
	if err != nil {
		fmt.Println("Error occured inside CreateFood in repo")
		fmt.Println(err.Error())
		return err
	}
	return err
}

func (e *EatRepository) UpdateEat(ctx context.Context, eat Eat) (string, error) {

	id, _ := strconv.Atoi(eat.Id)
	Foodid, _ := strconv.Atoi(eat.Foodid)

	fmt.Println(id)
	fmt.Println(Foodid)

	q := `UPDATE eat SET food_id=$1, name=$2, amount=$3, unit=$4, carb=$5, portein=$6, fat=$7, cal=$8 WHERE id=$9`
	res, err := e.db.ExecContext(ctx, q, Foodid, eat.Name, eat.Amount, eat.Unit, eat.Cal, eat.Portein, eat.Fat, eat.Cal, id)

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

func (e *EatRepository) DeleteEat(ctx context.Context, id int) (string, error) {
	res, err := e.db.ExecContext(ctx, "DELETE FROM eat WHERE id = $1 ", id)
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
