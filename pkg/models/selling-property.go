package models

import (
	"fmt"
	"time"

	"github.com/deemakuzovkin/investment-game/pkg/cached"
	"github.com/deemakuzovkin/investment-game/pkg/common"
	"github.com/deemakuzovkin/investment-game/pkg/utils"
	"github.com/google/uuid"
)

type SellingProperty struct {
	InvestorId string    `json:"investorId"`
	PropertyId string    `json:"propertyId"`
	Time       time.Time `json:"time"`
}

func (sell *SellingProperty) key() []byte {
	dataTime := sell.Time.Format("2006-01-02-15-04-05")
	fmt.Sprintf("%s-%s-%s", sell.InvestorId, sell.PropertyId, dataTime)

	return []byte(sell.InvestorId)
}

// AddOrUpdate add or update in database

func (sell *SellingProperty) AddOrUpdate() error {
	sellingPropertyBase, err := cached.Connect(common.TableInvestorProperties)
	if err != nil {
		return err
	}
	sellingPropertyData := utils.ToJsonBytes(sell)
	sellingPropertyKey := sell.key()
	return sellingPropertyBase.Add(sellingPropertyKey, sellingPropertyData)

}

// Sell property

func (sell *SellingProperty) Sell() error {
	investorData := Investor{
		Id: sell.InvestorId,
	}

	investor, err := investorData.Get()
	if err != nil {
		return err
	}

	investorId := investor.Id
	propertyData := Property{
		Id: sell.PropertyId,
	}

	property, err := propertyData.Get()
	if err != nil {
		return err
	}

	if sell.PropertyId == "" {
		return fmt.Errorf("Property not found")
	}

	transaction := BudgetTransaction{
		Id:         uuid.New().String(),
		InvestorId: investorId,
		Time:       time.Now(),
		Type:       common.SellingPropertyType,
		Sum:        property.Price,
	}

	return transaction.AddOrUpdate()
}
