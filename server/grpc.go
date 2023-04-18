package server

import (
	"database/sql"
	"fmt"
	bpb "go-kit-demo/body/proto"
	b "go-kit-demo/body/service"
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

	foodendpoint := f.MakeGrpcEndpoints(food)
	fgrpcServer := f.NewGRPCServer(foodendpoint, logger)
	bodypoint := b.MakeGrpcEndpoints(body)
	bgrpcServer := b.NewGRPCServer(bodypoint, logger)

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
		level.Info(logger).Log("msg", "Server started successfully 🚀")
		baseServer.Serve(grpcListener)
	}()

	level.Error(logger).Log("exit", <-errs)
}
