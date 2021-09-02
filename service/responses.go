package service

import "github.com/mrNobody95/mabna/models"

type ReportResponse struct {
	Report []models.Trade
	Error  error
}
