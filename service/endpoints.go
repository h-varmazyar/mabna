package service

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	GetTradeReport endpoint.Endpoint
}

func MakeEndpoints(svc Service, middlewares []endpoint.Middleware) Endpoints {
	return Endpoints{
		GetTradeReport: wrapEndpoint(makeGetReportEndpoint(svc), middlewares),
	}
}

func wrapEndpoint(e endpoint.Endpoint, middlewares []endpoint.Middleware) endpoint.Endpoint {
	for _, m := range middlewares {
		e = m(e)
	}
	return e
}

func makeGetReportEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		resp, err := svc.TradeReport(ctx)
		return &ReportResponse{
			Report: resp,
			Error: err,
		}, err
	}
}