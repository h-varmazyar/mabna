package service

import (
	"context"
	"github.com/go-kit/kit/log"
	"github.com/mrNobody95/mabna/models"
)

type Service interface {
	TradeReport(context.Context) ([]models.Trade, error)
}

type tradeService struct {
	log log.Logger
}

func NewService() Service {
	var svc Service
	{
		svc = &tradeService{}
	}
	return svc
}

func (service *tradeService) TradeReport(_ context.Context) ([]models.Trade, error) {
	return models.FetchLastInstrumentsTrade()
}