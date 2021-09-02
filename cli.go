package main

import (
	"fmt"
	"github.com/mrNobody95/mabna/models"
	log "github.com/sirupsen/logrus"
	"sync"
	"time"
)

func Cli() {
	for {
		instrumentCount := 0
		tradeCount := 0
		mapLock := sync.Mutex{}
		insDateMap := make(map[int][]time.Time)
		insId := make([]int, 0)
		wg := sync.WaitGroup{}

		fmt.Print("Please enter instrument count(type -1 for exit from data entrance):")
		_, err := fmt.Scan(&instrumentCount)
		if err != nil {
			log.Fatal(err)
		}
		if instrumentCount==-1 {
			fmt.Println("cli worker end")
			return
		}

		fmt.Print("Please enter trade count(type -1 for exit from data entrance):")
		_, err = fmt.Scan(&tradeCount)
		if err != nil {
			log.Fatal(err)
		}
		if tradeCount==-1 {
			fmt.Println("cli worker end")
			return
		}

		wg.Add(instrumentCount)
		for i := 0; i < instrumentCount; i++ {
			go func() {
				ins := models.CreateRandomInstrument()
				mapLock.Lock()
				insDateMap[ins.Id] = make([]time.Time, 0)
				insId = append(insId, ins.Id)
				mapLock.Unlock()
				wg.Done()
			}()
		}
		wg.Wait()

		wg.Add(tradeCount)
		for i := 0; i < tradeCount; i++ {
			go func() {
				id := insId[models.RandomInt(0, len(insId))]
				var date time.Time
			OUTER:
				for {
					time.Sleep(time.Second * 5)
					date = models.RandomDate(time.Now().Add(time.Hour*354*24*2*-1).Unix(), time.Now().Unix())
					if len(insDateMap[id]) == 0 {
						break OUTER
					}
					for _, d := range insDateMap[id] {
						fmt.Println(d, date)
						if d != date {
							break OUTER
						}
					}
				}
				models.CreateRandomTrade(id, date)
				wg.Done()
			}()
		}
		wg.Wait()
		fmt.Printf("%d trades craeted for %d instruments\n\n", tradeCount, instrumentCount)

	}
}
