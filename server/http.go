package server

import (
	"database/sql"
	"net/http"
	"os"

	b "go-kit-demo/body/service"
	e "go-kit-demo/eat/service"
	f "go-kit-demo/food/service"

	httptransport "github.com/go-kit/kit/transport/http"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/gorilla/mux"
)

func HttpRun(db *sql.DB, logger log.Logger) {
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
	DeleteBodyHandler := httptransport.NewServer(
		b.DeleteBodyEndpoint(body),
		b.DecodeDeleteBodyRequest,
		b.EncodeResponse,
	)

	r := mux.NewRouter()
	http.Handle("/", r)

	r.Handle("/food", CreateFoodHandler).Methods("POST")
	r.Handle("/food/{foodid:[0-9]+}", UpdateFoodHandler).Methods("PATCH")
	r.Handle("/food", GetAllFoodsHandler).Methods("GET")
	r.Handle("/food/{foodid:[0-9]+}", GetByFoodIdHandler).Methods("GET")
	r.Handle("/food/{foodid}", DeleteFoodHandler).Methods("DELETE")

	r.Handle("/eat", CreateEatHandler).Methods("POST")
	r.Handle("/eat/{eatid:[0-9]+}", UpdateEatHandler).Methods("PATCH")
	r.Handle("/eats", GetAllEatsHandler).Methods("GET")
	r.Handle("/eat/{eatid:[0-9]+}", GetByEatIdHandler).Methods("GET")
	r.Handle("/eat/{eatid:[0-9]+}", DeleteEatHandler).Methods("DELETE")

	r.Handle("/body", CreateBodyHandler).Methods("POST")
	r.Handle("/bodys", GetAllBodysHandler).Methods("GET")
	r.Handle("/body/{bodyid:[0-9]+}", GetByBodyIdHandler).Methods("GET")
	r.Handle("/body/{bodyid:[0-9]+}", DeleteBodyHandler).Methods("DELETE")

	http.ListenAndServe(":8787", nil)
}
