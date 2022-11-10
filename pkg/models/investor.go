package models

import (
	"encoding/json"
	"github.com/deemakuzovkin/investment-game/pkg/cached"
	"github.com/deemakuzovkin/investment-game/pkg/common"
	"github.com/deemakuzovkin/investment-game/pkg/utils"
)

// Investor type
type Investor struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Age     int    `json:"age"`
}

func (inv *Investor) key() []byte {
	return []byte(inv.Id)
}

// AddOrUpdate add or update in database
func (inv *Investor) AddOrUpdate() error {
	dataBase, err := cached.Connect(common.TableInvestor)
	if err != nil {
		return err
	}
	dataBytes := utils.ToJsonBytes(inv)
	key := inv.key()
	return dataBase.Add(key, dataBytes)
}

// Get  investor by key
func (inv *Investor) Get() (*Investor, error) {
	dataBase, err := cached.Connect(common.TableInvestor)
	if err != nil {
		return nil, err
	}
	key := inv.key()
	getByte, err := dataBase.Get(key)
	if err != nil {
		return nil, err
	}
	var result *Investor
	err = json.Unmarshal(getByte, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
