package models

import (
	"fmt"
	"github.com/deemakuzovkin/investment-game/pkg/cached"
	"github.com/deemakuzovkin/investment-game/pkg/common"
	"github.com/deemakuzovkin/investment-game/pkg/utils"
	"time"
)

type BudgetTransaction struct {
	Id         string    `json:"Id"`
	InvestorId string    `json:"investorId"`
	Time       time.Time `json:"time"`
	Type       string    `json:"type"`
	Sum        float64   `json:"sum"`
}

func (transaction *BudgetTransaction) key() []byte {
	dataTime := transaction.Time.Format("2006-01-02-15-04-05")
	key := fmt.Sprintf("%s-%s-%s", transaction.InvestorId, dataTime, transaction.Type)
	return []byte(key)
}

// AddOrUpdate add or update in database
func (transaction *BudgetTransaction) AddOrUpdate() error {
	transactionBase, err := cached.Connect(common.TableTransaction)
	if err != nil {
		return err
	}
	transactionData := utils.ToJsonBytes(transaction)
	transactionKey := transaction.key()
	err = transactionBase.Add(transactionKey, transactionData)
	if err != nil {
		return err
	}
	newBudget := Budget{
		InvestorId: transaction.InvestorId,
		Sum:        0,
	}
	currentBudget, err := newBudget.Get()
	if err != nil {
		newBudget.Sum = transaction.Sum
	} else {
		newBudget.Sum = currentBudget.Sum + transaction.Sum
	}
	return newBudget.AddOrUpdate()
}
