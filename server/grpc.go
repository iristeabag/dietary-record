package server

import (
	"database/sql"
	"fmt"
	bep "go-kit-demo/body/endpoint/grpc"
	bpb "go-kit-demo/body/proto"
	b "go-kit-demo/body/service"
	btp "go-kit-demo/body/transport/grpc"
	eep "go-kit-demo/eat/endpoint/grpc"
	epb "go-kit-demo/eat/proto"
	e "go-kit-demo/eat/service"
	etp "go-kit-demo/eat/transport/grpc"
	fpb "go-kit-demo/food/proto"
	f "go-kit-demo/food/service"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"google.golang.org/grpc"
)

func GrpcRun(db *sql.DB, logger log.Logger) {
	logger = log.NewJSONLogger(os.Stdout)
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	logger = log.With(logger, "caller", log.DefaultCaller)

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

	foodendpoint := f.MakeGrpcEndpoints(food)
	fgrpcServer := f.NewGRPCServer(foodendpoint, logger)
	bodypoint := bep.MakeGrpcEndpoints(body)
	bgrpcServer := btp.NewGRPCServer(bodypoint, logger)
	eatpoint := eep.MakeGrpcEndpoints(eat)
	egrpcServer := etp.NewGRPCServer(eatpoint, logger)

	errs := make(chan error)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGALRM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	grpcListener, err := net.Listen("tcp", ":50056")
	if err != nil {
		logger.Log("during", "Listen", "err", err)
		os.Exit(1)
	}

	go func() {
		baseServer := grpc.NewServer()
		fpb.RegisterFoodServiceServer(baseServer, fgrpcServer)
		bpb.RegisterBodyServiceServer(baseServer, bgrpcServer)
		epb.RegisterEatServiceServer(baseServer, egrpcServer)
		level.Info(logger).Log("msg", "Server started successfully 🚀")
		baseServer.Serve(grpcListener)
	}()

	level.Error(logger).Log("exit", <-errs)
}
