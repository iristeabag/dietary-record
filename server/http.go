package server

import (
	"database/sql"
	"net/http"
	"os"

	bep "go-kit-demo/body/endpoint/http"
	b "go-kit-demo/body/service"
	btp "go-kit-demo/body/transport/http"
	eep "go-kit-demo/eat/endpoint/http"
	e "go-kit-demo/eat/service"
	etp "go-kit-demo/eat/transport/http"
	fep "go-kit-demo/food/endpoint/http"
	f "go-kit-demo/food/service"
	ftp "go-kit-demo/food/transport/http"

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
		fep.CreateFoodEndpoint(food),
		ftp.DecodeCreateFoodRequest,
		ftp.EncodeResponse,
	)
	GetByFoodIdHandler := httptransport.NewServer(
		fep.GetFoodByIdEndpoint(food),
		ftp.DecodeGetFoodByIdRequest,
		ftp.EncodeResponse,
	)
	GetAllFoodsHandler := httptransport.NewServer(
		fep.GetAllFoodsEndpoint(food),
		ftp.DecodeGetAllFoodsRequest,
		ftp.EncodeResponse,
	)
	UpdateFoodHandler := httptransport.NewServer(
		fep.UpdateFoodEndpoint(food),
		ftp.DecodeUpdateFoodRequest,
		ftp.EncodeResponse,
	)
	DeleteFoodHandler := httptransport.NewServer(
		fep.DeleteFoodEndpoint(food),
		ftp.DecodeDeleteFoodRequest,
		ftp.EncodeResponse,
	)
	// Eat Handler
	CreateEatHandler := httptransport.NewServer(
		eep.CreateEatEndpoint(eat),
		etp.DecodeCreateEatRequest,
		etp.EncodeResponse,
	)
	GetByEatIdHandler := httptransport.NewServer(
		eep.GetEatByIdEndpoint(eat),
		etp.DecodeGetEatByIdRequest,
		etp.EncodeResponse,
	)
	GetAllEatsHandler := httptransport.NewServer(
		eep.GetAllEatsEndpoint(eat),
		etp.DecodeGetAllEatsRequest,
		etp.EncodeResponse,
	)
	UpdateEatHandler := httptransport.NewServer(
		eep.UpdateEatEndpoint(eat),
		etp.DecodeUpdateEatRequest,
		etp.EncodeResponse,
	)
	DeleteEatHandler := httptransport.NewServer(
		eep.DeleteEatEndpoint(eat),
		etp.DecodeDeleteEatRequest,
		etp.EncodeResponse,
	)
	// Body Handler
	CreateBodyHandler := httptransport.NewServer(
		bep.CreateBodyEndpoint(body),
		btp.DecodeCreateBodyRequest,
		btp.EncodeResponse,
	)
	GetByBodyIdHandler := httptransport.NewServer(
		bep.GetBodyByIdEndpoint(body),
		btp.DecodeGetBodyByIdRequest,
		btp.EncodeResponse,
	)
	GetAllBodysHandler := httptransport.NewServer(
		bep.GetAllBodysEndpoint(body),
		btp.DecodeGetAllBodysRequest,
		btp.EncodeResponse,
	)
	UpdateBodyHandler := httptransport.NewServer(
		bep.UpdateBodyEndpoint(body),
		btp.DecodeUpdateBodyRequest,
		btp.EncodeResponse,
	)
	DeleteBodyHandler := httptransport.NewServer(
		bep.DeleteBodyEndpoint(body),
		btp.DecodeDeleteBodyRequest,
		btp.EncodeResponse,
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
	r.Handle("/body/{bodyid:[0-9]+}", UpdateBodyHandler).Methods("PATCH")
	r.Handle("/bodys", GetAllBodysHandler).Methods("GET")
	r.Handle("/body/{bodyid:[0-9]+}", GetByBodyIdHandler).Methods("GET")
	r.Handle("/body/{bodyid:[0-9]+}", DeleteBodyHandler).Methods("DELETE")

	http.ListenAndServe(":8787", nil)
}
