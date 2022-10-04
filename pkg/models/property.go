package models

import (
	"encoding/json"
	"github.com/deemakuzovkin/investment-game/pkg/cached"
	"github.com/deemakuzovkin/investment-game/pkg/common"
	"github.com/deemakuzovkin/investment-game/pkg/utils"
)

// Property type
type Property struct {
	Id      string  `json:"id"`
	Name    string  `json:"name"`
	Comment string  `json:"comment"`
	Price   float64 `json:"price"`
}

func (prop *Property) key() []byte {
	return []byte(prop.Id)
}

// AddOrUpdate add or update in database
func (prop *Property) AddOrUpdate() error {
	dataBase, err := cached.Connect(common.TableProperty)
	if err != nil {
		return err
	}
	dataBytes := utils.ToJsonBytes(prop)
	key := prop.key()
	return dataBase.Add(key, dataBytes)
}

// Get property by key
func (prop *Property) Get() (*Property, error) {
	dataBase, err := cached.Connect(common.TableProperty)
	if err != nil {
		return nil, err
	}
	key := prop.key()
	getByte, err := dataBase.Get(key)
	if err != nil {
		return nil, err
	}
	var result *Property
	err = json.Unmarshal(getByte, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
