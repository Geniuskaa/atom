package errs

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
)

type Error struct {
	httpStatus int
	grpcStatus codes.Code
	msg        string
}

func (e *Error) Error() string {
	return e.msg
}

// It returns Error struct with appropriate in meaning gRPC response status code base on https status code.
func NewError(httpStatus int, msg string) *Error {
	e := Error{
		httpStatus: httpStatus,
		msg:        msg,
	}

	switch httpStatus {
	case http.StatusOK:
		e.grpcStatus = codes.OK
	case 499:
		// HTTP 499 Client Closed Request
		e.grpcStatus = codes.Canceled
	case http.StatusBadRequest:
		e.grpcStatus = codes.InvalidArgument
	case http.StatusGatewayTimeout:
		e.grpcStatus = codes.DeadlineExceeded
	case http.StatusNotFound:
		e.grpcStatus = codes.NotFound
	case http.StatusConflict:
		e.grpcStatus = codes.AlreadyExists
	case http.StatusForbidden:
		e.grpcStatus = codes.PermissionDenied
	case http.StatusTooManyRequests:
		e.grpcStatus = codes.ResourceExhausted
	case http.StatusNotImplemented:
		e.grpcStatus = codes.Unimplemented
	case http.StatusServiceUnavailable:
		e.grpcStatus = codes.Unavailable
	case http.StatusInternalServerError:
		e.grpcStatus = codes.Internal
	case http.StatusUnauthorized:
		e.grpcStatus = codes.Unauthenticated

	default:
		e.grpcStatus = codes.Unknown
	}

	return &e
}

func NewErrorWithGrpcStatus(httpStatus int, grpcStatus codes.Code, msg string) *Error {
	e := Error{
		httpStatus: httpStatus,
		grpcStatus: grpcStatus,
		msg:        msg,
	}

	return &e
}

// Returns http status code.
func (e *Error) StatusCode() int {
	return e.httpStatus
}

// Returns msg that will be shown API client. Msg can be empty.
func (e *Error) Msg() string {
	return e.msg
}

// Returns gRPC error for API client.
func (e *Error) GrpcErr() error {
	return status.Error(e.grpcStatus, e.msg)
}
