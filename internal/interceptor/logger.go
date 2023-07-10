package interceptor

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type Logger struct {
	logger *zap.SugaredLogger
}

func NewLogger(logger *zap.SugaredLogger) *Logger {
	return &Logger{logger: logger}
}

// Logging from URL:
// https://dev.to/techschoolguru/use-grpc-interceptor-for-authorization-with-jwt-1c5h
func (l *Logger) Unary() grpc.UnaryServerInterceptor {
	return func(ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler) (interface{}, error) {

		err := l.log(ctx, info.FullMethod)
		if err != nil {
			l.logger.Errorf("Unary logger failed: %w", err)
		}

		return handler(ctx, req)
	}
}

// Logging from URL:
// https://qaa-engineer.ru/kak-vypolnyat-logirovanie-raboty-grpc/
func (l *Logger) Unary_V2() grpc.UnaryServerInterceptor {
	return func(ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler) (interface{}, error) {

		l.logger.Info("RPC call",
			zap.String("method", info.FullMethod),
			zap.Any("request", req),
		)

		resp, err := handler(ctx, req)

		if err != nil {
			l.logger.Error("RPC error",
				zap.String("method", info.FullMethod),
				zap.Error(err),
			)
		} else {
			l.logger.Info("RPC response",
				zap.String("method", info.FullMethod),
				zap.Any("response", resp),
			)
		}

		return resp, err
	}
}

func (l *Logger) Stream() grpc.StreamServerInterceptor {
	return func(srv interface{},
		stream grpc.ServerStream,
		info *grpc.StreamServerInfo,
		handler grpc.StreamHandler) error {

		fmt.Println("Do some work...")
		fmt.Println("--> stream interceptor: ", info.FullMethod)

		return handler(srv, stream)
	}
}

func (l *Logger) log(ctx context.Context, method string) error {

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return status.Errorf(codes.Unauthenticated, "metadata is not provided")
	}

	l.logger.Infof("%s : %v", method, md)

	return nil
}
