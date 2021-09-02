package service

import (
	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func Start()  {
	var middlewares []endpoint.Middleware
	var options []httptransport.ServerOption
	svc := NewService()
	eps := MakeEndpoints(svc, middlewares)
	r := mux.NewRouter()
	r.Methods(http.MethodGet).Path("/report").Handler(GetTradeReportHandler(eps.GetTradeReport, options))
	log.Infof("Statring service on 127.0.0.1:8080\n")
	svr := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: r,
	}
	log.Errorf("serving failed: %v",svr.ListenAndServe())
}
