package utilities

import (
	"github.com/grpc-ecosystem/grpc-gateway/grpcgw"
)

type OpCode = grpcgw.OpCode

const (
	OpNop     = grpcgw.OpNop
	OpPush    = grpcgw.OpPush
	OpLitPush = grpcgw.OpLitPush
	OpPushM   = grpcgw.OpPushM
	OpConcatN = grpcgw.OpConcatN
	OpCapture = grpcgw.OpCapture
	OpEnd     = grpcgw.OpEnd
)

type DoubleArray = grpcgw.DoubleArray

var NewDoubleArray = grpcgw.NewDoubleArray
