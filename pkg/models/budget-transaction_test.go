package models

import (
	"github.com/deemakuzovkin/investment-game/pkg/common"
	"github.com/google/uuid"
	"reflect"
	"testing"
	"time"
)

func TestBudgetTransaction_AddOrUpdate(t *testing.T) {
	testInvestorId := uuid.New().String()
	testBudget := Budget{
		InvestorId: testInvestorId,
		Sum:        1200.0,
	}
	type fields struct {
		Id            string
		InvestorId    string
		Time          time.Time
		Type          string
		Sum           float64
		CurrentBudget *Budget
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
		want    float64
	}{
		{
			name: "add new budget",
			fields: fields{
				Id:            uuid.New().String(),
				InvestorId:    testInvestorId,
				Time:          time.Now(),
				Type:          common.AddBudgetTransactionType,
				Sum:           200,
				CurrentBudget: nil,
			},
			wantErr: false,
			want:    200,
		},
		{
			name: "add current budget",
			fields: fields{
				Id:            uuid.New().String(),
				InvestorId:    testInvestorId,
				Time:          time.Now(),
				Type:          common.AddBudgetTransactionType,
				Sum:           200,
				CurrentBudget: &testBudget,
			},
			wantErr: false,
			want:    1400,
		},
		{
			name: "add current budget",
			fields: fields{
				Id:            uuid.New().String(),
				InvestorId:    testInvestorId,
				Time:          time.Now(),
				Type:          common.AddBudgetTransactionType,
				Sum:           400,
				CurrentBudget: &testBudget,
			},
			wantErr: false,
			want:    1600,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transaction := &BudgetTransaction{
				Id:         tt.fields.Id,
				InvestorId: tt.fields.InvestorId,
				Time:       tt.fields.Time,
				Type:       tt.fields.Type,
				Sum:        tt.fields.Sum,
			}
			totalBudget := &Budget{
				InvestorId: transaction.InvestorId,
				Sum:        tt.fields.Sum,
			}
			if tt.fields.CurrentBudget != nil {
				err := tt.fields.CurrentBudget.AddOrUpdate()
				if err != nil {
					t.Errorf("Budget AddOrUpdate() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				totalBudget = tt.fields.CurrentBudget
			}
			if err := transaction.AddOrUpdate(); (err != nil) != tt.wantErr {
				t.Errorf("AddOrUpdate() error = %v, wantErr %v", err, tt.wantErr)
			}
			equalBudget, err := totalBudget.Get()
			if err != nil {
				t.Errorf("Get total budget error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(equalBudget.Sum, tt.want) {
				t.Errorf("Budget equals got = %v, want %v", totalBudget.Sum, tt.want)
			}
		})
	}
}
