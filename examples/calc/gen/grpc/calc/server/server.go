package server

import (
	"context"

	calcsvc "goa.design/goa/examples/calc/gen/calc"
	calcpb "goa.design/goa/examples/calc/gen/grpc/calc"
)

type Server struct {
	endpoints *calcsvc.Endpoints
}

func New(e *calcsvc.Endpoints) *Server {
	return &Server{
		endpoints: e,
	}
}

func (s *Server) Add(ctx context.Context, p *calcpb.AddPayload) (*calcpb.AddResponse, error) {
	payload := &calcsvc.AddPayload{
		A: int(p.GetA()),
		B: int(p.GetB()),
	}
	res, err := s.endpoints.Add(ctx, payload)
	resInt := res.(int)
	return &calcpb.AddResponse{
		AddResponseField: int32(resInt),
	}, err
}
