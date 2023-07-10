package main

import (
	"atom/internal/config"
	"atom/internal/database"
	"atom/internal/server"
	"context"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"golang.org/x/sync/errgroup"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	mainCtx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGHUP, syscall.SIGINT,
		syscall.SIGQUIT, syscall.SIGTERM)
	defer stop()

	logger, _ := loggerInit()

	conf, err := config.NewConfig()
	if err != nil {
		logger.Panic(err)
	}

	pool, err := database.NewPostgresPool(mainCtx, conf)
	if err != nil {
		logger.Panic(err)
	}

	grpcServer, err := server.NewGrpcServer(logger, conf)
	if err != nil {
		logger.Panic(err)
	}

	err = grpcServer.InitServices(pool)
	if err != nil {
		logger.Panic(err)
	}

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", conf.App.Addr, conf.App.Port))
	if err != nil {
		logger.Panic(err)
	}

	g, gCtx := errgroup.WithContext(mainCtx)
	g.Go(func() error {
		fmt.Println("Grpc server started!")
		return grpcServer.Serve(listener)
	})
	g.Go(func() error {
		<-gCtx.Done()

		fmt.Println("Grpc server began graceful shut down.")
		grpcServer.GracefulStop()
		fmt.Println("Grpc server graceful shut down.")
		return nil
	})
	g.Go(func() error {

		<-gCtx.Done()
		pool.Close()
		return nil
	})

	if err := g.Wait(); err != nil {
		fmt.Printf("exit reason: %s \n", err)
	}

	fmt.Println("Server was gracefully shut down.")
}

func loggerInit() (*zap.SugaredLogger, zap.AtomicLevel) {
	encoderConfig := zap.NewDevelopmentEncoderConfig()
	encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC1123Z)
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	fileEncoder := zapcore.NewJSONEncoder(encoderConfig)
	consoleEncoder := zapcore.NewConsoleEncoder(encoderConfig)

	file, err := os.OpenFile("./logs/logs.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666) // os.Create("./logs/logs.txt")

	if err != nil {
		panic("Error with creating or opening file")
	}

	writeSyncer := zapcore.AddSync(file)
	atom := zap.NewAtomicLevelAt(zapcore.InfoLevel)
	core := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, writeSyncer, atom),
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), atom),
	)

	sugarLogger := zap.New(core).Sugar()

	return sugarLogger, atom
}
