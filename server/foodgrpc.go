package server

import (
	"database/sql"
	"fmt"
	fep "go-kit-demo/food/endpoint/grpc"
	f "go-kit-demo/food/service"
	ftp "go-kit-demo/food/transport/grpc"
	fpb "go-kit-demo/proto/food"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"google.golang.org/grpc"
)

func FoodGrpcRun(db *sql.DB, logger log.Logger) {
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

	foodendpoint := fep.MakeGrpcEndpoints(food)
	fgrpcServer := ftp.NewGRPCServer(foodendpoint, logger)

	errs := make(chan error)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGALRM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	grpcListener, err := net.Listen("tcp", ":50057")
	if err != nil {
		logger.Log("during", "Listen", "err", err)
		os.Exit(1)
	}

	go func() {
		baseServer := grpc.NewServer()
		fpb.RegisterFoodServiceServer(baseServer, fgrpcServer)
		level.Info(logger).Log("msg", "Food Server started successfully 🚀")
		baseServer.Serve(grpcListener)
	}()

	level.Error(logger).Log("exit", <-errs)
}
