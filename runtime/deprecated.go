package runtime

import "github.com/grpc-ecosystem/grpc-gateway/grpcgw"

// Types exported from package grpcgw
type (
	Decoder               = grpcgw.Decoder
	DecoderFunc           = grpcgw.DecoderFunc
	Delimited             = grpcgw.Delimited
	Encoder               = grpcgw.Encoder
	EncoderFunc           = grpcgw.EncoderFunc
	HandlerFunc           = grpcgw.HandlerFunc
	HeaderMatcherFunc     = grpcgw.HeaderMatcherFunc
	JSONBuiltin           = grpcgw.JSONBuiltin
	JSONPb                = grpcgw.JSONPb
	Marshaler             = grpcgw.Marshaler
	Pattern               = grpcgw.Pattern
	ProtoErrorHandlerFunc = grpcgw.ProtoErrorHandlerFunc
	ProtoMarshaller       = grpcgw.ProtoMarshaller
	ServeMux              = grpcgw.ServeMux
	ServeMuxOption        = grpcgw.ServeMuxOption
	ServerMetadata        = grpcgw.ServerMetadata
)

// Package-level vars exported from package grpcgw
var (
	HTTPError             = grpcgw.HTTPError
	ErrNotMatch           = grpcgw.ErrNotMatch
	DefaultContextTimeout = grpcgw.DefaultContextTimeout
)

// Package-level funcs exported from package grpcgw
var (
	AnnotateContext              = grpcgw.AnnotateContext
	Bool                         = grpcgw.Bool
	BoolP                        = grpcgw.BoolP
	Bytes                        = grpcgw.Bytes
	DefaultHTTPError             = grpcgw.DefaultHTTPError
	DefaultHTTPProtoErrorHandler = grpcgw.DefaultHTTPProtoErrorHandler
	DefaultHeaderMatcher         = grpcgw.DefaultHeaderMatcher
	DefaultOtherErrorHandler     = grpcgw.DefaultOtherErrorHandler
	Duration                     = grpcgw.Duration
	Float32                      = grpcgw.Float32
	Float32P                     = grpcgw.Float32P
	Float64                      = grpcgw.Float64
	Float64P                     = grpcgw.Float64P
	ForwardResponseMessage       = grpcgw.ForwardResponseMessage
	ForwardResponseStream        = grpcgw.ForwardResponseStream
	HTTPStatusFromCode           = grpcgw.HTTPStatusFromCode
	Int32                        = grpcgw.Int32
	Int32P                       = grpcgw.Int32P
	Int64                        = grpcgw.Int64
	Int64P                       = grpcgw.Int64P
	NewServerMetadataContext     = grpcgw.NewServerMetadataContext
	PopulateFieldFromPath        = grpcgw.PopulateFieldFromPath
	PopulateQueryParameters      = grpcgw.PopulateQueryParameters
	String                       = grpcgw.String
	StringP                      = grpcgw.StringP
	Timestamp                    = grpcgw.Timestamp
	Uint32                       = grpcgw.Uint32
	Uint32P                      = grpcgw.Uint32P
	Uint64                       = grpcgw.Uint64
	Uint64P                      = grpcgw.Uint64P
)
