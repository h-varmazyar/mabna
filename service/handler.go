package service

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

func GetTradeReportHandler(ep endpoint.Endpoint, options []httptransport.ServerOption) *httptransport.Server {
	return httptransport.NewServer(
		ep,
		decodeGetTradeReportRequest,
		encodeGetTradeReportResponse,
		options...,
	)
}

func decodeGetTradeReportRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req ReportRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}

func encodeGetTradeReportResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	resp, ok := response.(*ReportResponse)
	if !ok {
		return errors.New("error decoding")
	}
	return json.NewEncoder(w).Encode(resp)
}
