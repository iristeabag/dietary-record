package server

import (
	"database/sql"
	"fmt"
	bep "go-kit-demo/body/endpoint/grpc"
	b "go-kit-demo/body/service"
	btp "go-kit-demo/body/transport/grpc"
	eep "go-kit-demo/eat/endpoint/grpc"
	e "go-kit-demo/eat/service"
	etp "go-kit-demo/eat/transport/grpc"
	bpb "go-kit-demo/proto/body"
	epb "go-kit-demo/proto/eat"
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

	grpcListener, err := net.Listen("tcp", ":50058")
	if err != nil {
		logger.Log("during", "Listen", "err", err)
		os.Exit(1)
	}

	go func() {
		baseServer := grpc.NewServer()
		bpb.RegisterBodyServiceServer(baseServer, bgrpcServer)
		epb.RegisterEatServiceServer(baseServer, egrpcServer)
		level.Info(logger).Log("msg", "Server started successfully 🚀")
		baseServer.Serve(grpcListener)
	}()

	level.Error(logger).Log("exit", <-errs)
}
