package models

import (
	"time"
)

type Trade struct {
	Id           int
	InstrumentId int        `gorm:"primaryKey"`
	Instrument   Instrument `gorm:"foreignKey:InstrumentId"`
	DateEn       time.Time  `gorm:"primaryKey"`
	Open         float64
	High         float64
	Low          float64
	Close        float64
}

func FetchLastInstrumentsTrade() ([]Trade, error) {
	var trades []Trade
	err := db.Raw(`
SELECT i.*, t.*
FROM instruments i INNER JOIN
  (
     SELECT instrument_id,
     MAX(date_en) MaxDate
     FROM trades
     GROUP BY instrument_id
  ) MaxDates ON i.id = MaxDates.instrument_id INNER JOIN
    trades t ON MaxDates.instrument_id = t.instrument_id
    AND MaxDates.MaxDate = t.date_en
`).Preload("Instrument").Find(&trades).Error
	return trades, err
}

func CreateRandomTrade(instrumentId int, date time.Time) {
	(&Trade{
		Id:           1,
		InstrumentId: instrumentId,
		DateEn:       date,
		Open:         RandomFloat64(0, 2000),
		High:         RandomFloat64(0, 2000),
		Low:          RandomFloat64(0, 2000),
		Close:        RandomFloat64(0, 2000),
	}).Create()
}

func (trade *Trade) Create() {
	db.Model(&Trade{}).Create(&trade)
}
