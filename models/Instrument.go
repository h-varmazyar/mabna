package models

type Instrument struct {
	Id   int    `gorm:"primarykey"`
	Name string `gorm:"size:255;name"`
}

func CreateRandomInstrument() *Instrument{
	ins:=Instrument{
		Name: RandomString(5, "ABCDEFGHIJKLMNOPQRSTUVWXYZ"),
	}
	ins.Create()
	return &ins
}

func (instrument *Instrument) Create() {
	db.Model(&Instrument{}).Create(&instrument)
}