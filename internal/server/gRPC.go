package server

import (
	"atom/internal/config"
	"atom/internal/database/repository"
	"atom/internal/interceptor"
	"atom/internal/service/fleetManagmentSystem/saeJ1939"
	carSrvGrpcV1 "atom/internal/service/generated/car/v1"
	"crypto/tls"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"time"
)

type GRPC struct {
	logger *zap.SugaredLogger
	*grpc.Server
}

func loadTLSCredentials() (credentials.TransportCredentials, error) {
	// Load server's certificate and private key
	serverCert, err := tls.LoadX509KeyPair("certs/server-cert.pem", "certs/server-key.pem")
	if err != nil {
		return nil, err
	}

	// Create the credentials and return it
	conf := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientAuth:   tls.NoClientCert,
	}

	return credentials.NewTLS(conf), nil
}

func NewGrpcServer(logger *zap.SugaredLogger, config *config.Entity) (*GRPC, error) {
	srv := GRPC{logger: logger}

	// Закидываем TLS сертификат.
	//creds, err := credentials.NewServerTLSFromFile("certs/server.crt", "certs/server.key")
	//if err != nil {
	//	return nil, fmt.Errorf("NewGrpcServer could not load TLS keys: %w", err)
	//}

	loggerInterceptor := interceptor.NewLogger(logger)

	tlsCreds, err := loadTLSCredentials()
	if err != nil {
		return nil, fmt.Errorf("NewGrpcServer could not load TLS keys: %w", err)
	}

	// Тут вписываем опции для сервера.
	opts := []grpc.ServerOption{
		grpc.Creds(tlsCreds),
		grpc.ConnectionTimeout(time.Minute),
		grpc.MaxConcurrentStreams(uint32(config.App.MaxConcurrentStreams)),
		grpc.UnaryInterceptor(loggerInterceptor.Unary()),
	}

	grpcServer := grpc.NewServer(opts...)
	srv.Server = grpcServer

	return &srv, nil
}

func (g *GRPC) InitServices(pool *pgxpool.Pool) error {

	j1939Srv := saeJ1939.NewService()

	carRepo := repository.NewCarRepo(pool)
	carSrvGrpcV1.RegisterCarServiceServer(g.Server, carSrvGrpcV1.NewService(g.logger, carRepo, j1939Srv))

	return nil
}
