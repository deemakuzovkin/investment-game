package models

import (
	"encoding/json"
	"github.com/deemakuzovkin/investment-game/pkg/cached"
	"github.com/deemakuzovkin/investment-game/pkg/common"
	"github.com/deemakuzovkin/investment-game/pkg/utils"
	"time"
)

// DayEvent day event
type DayEvent struct {
	Id      string    `json:"id"`
	Time    time.Time `json:"time"`
	Name    string    `json:"name"`
	Comment string    `json:"comment"`
	Kf      float64   `json:"kf"`
}

func (prop *DayEvent) key() []byte {
	return []byte(prop.Id)
}

// AddOrUpdate add or update in database
func (prop *DayEvent) AddOrUpdate() error {
	dataBase, err := cached.Connect(common.TableDayEvent)
	if err != nil {
		return err
	}
	dataBytes := utils.ToJsonBytes(prop)
	key := prop.key()
	err = dataBase.Add(key, dataBytes)
	if err != nil {
		return err
	}
	lastDataBase, err := cached.Connect(common.TableLastDayEvent)
	if err != nil {
		return err
	}
	key = []byte(common.LatestKey)
	return lastDataBase.Add(key, dataBytes)
}

// Get day event by key
func (prop *DayEvent) Get() (*DayEvent, error) {
	dataBase, err := cached.Connect(common.TableDayEvent)
	if err != nil {
		return nil, err
	}
	key := prop.key()
	getByte, err := dataBase.Get(key)
	if err != nil {
		return nil, err
	}
	var result *DayEvent
	err = json.Unmarshal(getByte, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetLatest last day event
func (prop *DayEvent) GetLatest() (*DayEvent, error) {
	lastDataBase, err := cached.Connect(common.TableLastDayEvent)
	if err != nil {
		return nil, err
	}
	key := []byte(common.LatestKey)
	getByte, err := lastDataBase.Get(key)
	if err != nil {
		return nil, err
	}
	var result *DayEvent
	err = json.Unmarshal(getByte, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
