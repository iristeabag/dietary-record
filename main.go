package main

import (
	"database/sql"
	"fmt"
	b "go-kit-demo/body/service"
	e "go-kit-demo/eat/service"
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

	var body b.IBodyService
	body = b.BodyService{}
	{
		bodyRepo, err := b.NewBodyRepository(db, logger)
		if err != nil {
			level.Error(logger).Log("exit", err)
			os.Exit(-1)
		}
		body = b.NewBodyService(bodyRepo, logger)
	}

	var eat e.IEatService
	eat = e.EatService{}
	{
		eatRepo, err := e.NewEatRepository(db, logger)
		if err != nil {
			level.Error(logger).Log("exit", err)
			os.Exit(-1)
		}
		eat = e.NewEatService(eatRepo, logger)
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
	// Eat Handler
	CreateEatHandler := httptransport.NewServer(
		e.CreateEatEndpoint(eat),
		e.DecodeCreateEatRequest,
		e.EncodeResponse,
	)
	GetByEatIdHandler := httptransport.NewServer(
		e.GetEatByIdEndpoint(eat),
		e.DecodeGetEatByIdRequest,
		e.EncodeResponse,
	)
	GetAllEatsHandler := httptransport.NewServer(
		e.GetAllEatsEndpoint(eat),
		e.DecodeGetAllEatsRequest,
		e.EncodeResponse,
	)
	UpdateEatHandler := httptransport.NewServer(
		e.UpdateEatEndpoint(eat),
		e.DecodeUpdateEatRequest,
		e.EncodeResponse,
	)
	DeleteEatHandler := httptransport.NewServer(
		e.DeleteEatEndpoint(eat),
		e.DecodeDeleteEatRequest,
		e.EncodeResponse,
	)
	// Body Handler
	CreateBodyHandler := httptransport.NewServer(
		b.CreateBodyEndpoint(body),
		b.DecodeCreateBodyRequest,
		b.EncodeResponse,
	)
	GetByBodyIdHandler := httptransport.NewServer(
		b.GetBodyByIdEndpoint(body),
		b.DecodeGetBodyByIdRequest,
		b.EncodeResponse,
	)
	GetAllBodysHandler := httptransport.NewServer(
		b.GetAllBodysEndpoint(body),
		b.DecodeGetAllBodysRequest,
		b.EncodeResponse,
	)
	UpdateBodyHandler := httptransport.NewServer(
		b.UpdateBodyEndpoint(body),
		b.DecodeUpdateBodyRequest,
		b.EncodeResponse,
	)
	DeleteBodyHandler := httptransport.NewServer(
		b.DeleteBodyEndpoint(body),
		b.DecodeDeleteBodyRequest,
		b.EncodeResponse,
	)

	r := mux.NewRouter()
	http.Handle("/", r)
	http.Handle("/food", CreateFoodHandler)
	http.Handle("/food/update", UpdateFoodHandler)
	r.Handle("/food/all", GetAllFoodsHandler).Methods("GET")
	r.Handle("/food/{foodid}", GetByFoodIdHandler).Methods("GET")
	r.Handle("/food/{foodid}", DeleteFoodHandler).Methods("DELETE")

	http.Handle("/eat", CreateEatHandler)
	http.Handle("/eat/update", UpdateEatHandler)
	r.Handle("/eat/all", GetAllEatsHandler).Methods("GET")
	r.Handle("/eat/{eatid}", GetByEatIdHandler).Methods("GET")
	r.Handle("/eat/{eatid}", DeleteEatHandler).Methods("DELETE")

	http.Handle("/body", CreateBodyHandler)
	http.Handle("/body/update", UpdateBodyHandler)
	r.Handle("/body/all", GetAllBodysHandler).Methods("GET")
	r.Handle("/body/{bodyid}", GetByBodyIdHandler).Methods("GET")
	r.Handle("/body/{bodyid}", DeleteBodyHandler).Methods("DELETE")

	http.ListenAndServe(":8080", nil)
}
