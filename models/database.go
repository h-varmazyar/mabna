package models

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/mrNobody95/gorm"
	log "github.com/sirupsen/logrus"
	"os"
	"strconv"
)

var db *gorm.DB

func init() {
	var err error
	tmp := os.Getenv("DB_PORT")
	port, err := strconv.Atoi(tmp)
	if err != nil {
		log.Panicf("db port cast failed: %v", err)
	}
	err, db = (&gorm.DatabaseConfig{
		Type:     gorm.MYSQL,
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		CharSet:  os.Getenv("DB_CHARSET"),
		Name:     os.Getenv("db_name"),
		Host:     os.Getenv("DB_HOST"),
		Port:     port,
		SSLMode:  false,
	}).Initialize(&Instrument{}, &Trade{})

	if err != nil {
		log.Panicf("initializing db failed:%v", err)
	}
}
