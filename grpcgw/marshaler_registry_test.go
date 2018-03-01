package grpcgw_test

import (
	"errors"
	"io"
	"net/http"
	"testing"

	"github.com/grpc-ecosystem/grpc-gateway/grpcgw"
)

func TestMarshalerForRequest(t *testing.T) {
	r, err := http.NewRequest("GET", "http://example.com", nil)
	if err != nil {
		t.Fatalf(`http.NewRequest("GET", "http://example.com", nil) failed with %v; want success`, err)
	}
	r.Header.Set("Accept", "application/x-out")
	r.Header.Set("Content-Type", "application/x-in")

	mux := grpcgw.NewServeMux()

	in, out := grpcgw.MarshalerForRequest(mux, r)
	if _, ok := in.(*grpcgw.JSONPb); !ok {
		t.Errorf("in = %#v; want a grpcgw.JSONPb", in)
	}
	if _, ok := out.(*grpcgw.JSONPb); !ok {
		t.Errorf("out = %#v; want a grpcgw.JSONPb", in)
	}

	var marshalers [3]dummyMarshaler
	specs := []struct {
		opt grpcgw.ServeMuxOption

		wantIn  grpcgw.Marshaler
		wantOut grpcgw.Marshaler
	}{
		{
			opt:     grpcgw.WithMarshalerOption(grpcgw.MIMEWildcard, &marshalers[0]),
			wantIn:  &marshalers[0],
			wantOut: &marshalers[0],
		},
		{
			opt:     grpcgw.WithMarshalerOption("application/x-in", &marshalers[1]),
			wantIn:  &marshalers[1],
			wantOut: &marshalers[0],
		},
		{
			opt:     grpcgw.WithMarshalerOption("application/x-out", &marshalers[2]),
			wantIn:  &marshalers[1],
			wantOut: &marshalers[2],
		},
	}
	for i, spec := range specs {
		var opts []grpcgw.ServeMuxOption
		for _, s := range specs[:i+1] {
			opts = append(opts, s.opt)
		}
		mux = grpcgw.NewServeMux(opts...)

		in, out = grpcgw.MarshalerForRequest(mux, r)
		if got, want := in, spec.wantIn; got != want {
			t.Errorf("in = %#v; want %#v", got, want)
		}
		if got, want := out, spec.wantOut; got != want {
			t.Errorf("out = %#v; want %#v", got, want)
		}
	}

	r.Header.Set("Content-Type", "application/x-another")
	in, out = grpcgw.MarshalerForRequest(mux, r)
	if got, want := in, &marshalers[1]; got != want {
		t.Errorf("in = %#v; want %#v", got, want)
	}
	if got, want := out, &marshalers[0]; got != want {
		t.Errorf("out = %#v; want %#v", got, want)
	}
}

type dummyMarshaler struct{}

func (dummyMarshaler) ContentType() string { return "" }
func (dummyMarshaler) Marshal(interface{}) ([]byte, error) {
	return nil, errors.New("not implemented")
}

func (dummyMarshaler) Unmarshal([]byte, interface{}) error {
	return errors.New("not implemented")
}

func (dummyMarshaler) NewDecoder(r io.Reader) grpcgw.Decoder {
	return dummyDecoder{}
}
func (dummyMarshaler) NewEncoder(w io.Writer) grpcgw.Encoder {
	return dummyEncoder{}
}

type dummyDecoder struct{}

func (dummyDecoder) Decode(interface{}) error {
	return errors.New("not implemented")
}

type dummyEncoder struct{}

func (dummyEncoder) Encode(interface{}) error {
	return errors.New("not implemented")
}
