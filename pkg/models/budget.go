package models

import (
	"encoding/json"
	"github.com/deemakuzovkin/investment-game/pkg/cached"
	"github.com/deemakuzovkin/investment-game/pkg/common"
	"github.com/deemakuzovkin/investment-game/pkg/utils"
)

// Budget type
type Budget struct {
	InvestorId string  `json:"investorId"`
	Sum        float64 `json:"sum"`
}

func (budget *Budget) key() []byte {
	return []byte(budget.InvestorId)
}

// AddOrUpdate add or update in database
func (budget *Budget) AddOrUpdate() error {
	dataBase, err := cached.Connect(common.TableBudgets)
	if err != nil {
		return err
	}
	dataBytes := utils.ToJsonBytes(budget)
	key := budget.key()
	return dataBase.Add(key, dataBytes)
}

// Get property by key
func (budget *Budget) Get() (*Budget, error) {
	dataBase, err := cached.Connect(common.TableBudgets)
	if err != nil {
		return nil, err
	}
	key := budget.key()
	getByte, err := dataBase.Get(key)
	if err != nil {
		return nil, err
	}
	var result *Budget
	err = json.Unmarshal(getByte, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
