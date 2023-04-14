package main

import (
	"database/sql"
	"fmt"
	f "go-kit-demo/food/service"
	"net/http"
	"os"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 54320
	user     = "postgres"
	password = "postgres"
	dbname   = "dietary"
)

func GetDbConn() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	fmt.Println("Successfully connected!")
	return db
}

func main() {
	logger := log.NewLogfmtLogger(os.Stderr)
	db := GetDbConn()

	var food f.IFoodService
	food = f.FoodService{}
	{
		foodRepo, err := f.NewFoodRepository(db, logger)
		if err != nil {
			level.Error(logger).Log("exit", err)
			os.Exit(-1)
		}
		food = f.NewFoodService(foodRepo, logger)
	}

	// Food Handler
	CreateFoodHandler := httptransport.NewServer(
		f.CreateFoodEndpoint(food),
		f.DecodeCreateFoodRequest,
		f.EncodeResponse,
	)
	GetByFoodIdHandler := httptransport.NewServer(
		f.GetFoodByIdEndpoint(food),
		f.DecodeGetFoodByIdRequest,
		f.EncodeResponse,
	)
	GetAllFoodsHandler := httptransport.NewServer(
		f.GetAllFoodsEndpoint(food),
		f.DecodeGetAllFoodsRequest,
		f.EncodeResponse,
	)
	UpdateFoodHandler := httptransport.NewServer(
		f.UpdateFoodEndpoint(food),
		f.DecodeUpdateFoodRequest,
		f.EncodeResponse,
	)
	DeleteFoodHandler := httptransport.NewServer(
		f.DeleteFoodEndpoint(food),
		f.DecodeDeleteFoodRequest,
		f.EncodeResponse,
	)

	r := mux.NewRouter()
	http.Handle("/", r)
	http.Handle("/food", CreateFoodHandler)
	http.Handle("/food/update", UpdateFoodHandler)
	r.Handle("/food/all", GetAllFoodsHandler).Methods("GET")
	r.Handle("/food/{foodid}", GetByFoodIdHandler).Methods("GET")
	r.Handle("/food/{foodid}", DeleteFoodHandler).Methods("DELETE")

	http.ListenAndServe(":8080", nil)
}
