package models

import (
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestSellingProperty_AddOrUpdate(t *testing.T) {
	// Создаем тестового инвестора с бюджетом 1200
	testInvestorId := uuid.New().String()
	testBudget := Budget{
		InvestorId: testInvestorId,
		Sum:        1200.0,
	}
	testBudget.AddOrUpdate()

	// Создаем тестовую недвижимость с ценой 100
	testProperty := &Property{
		Id:      uuid.New().String(),
		Name:    "Property",
		Comment: "Some comment",
		Price:   100,
		Profit:  10,
	}
	testProperty.AddOrUpdate()

	type fields struct {
		InvestorId string
		PropertyId string
		Time       time.Time
	}

	tests := []struct {
		name    string
		fields  fields
		wantErr bool
		want    float64
	}{
		{
			name: "sell property",
			fields: fields{
				InvestorId: testInvestorId,
				PropertyId: testProperty.Id,
				Time:       time.Now(),
			},
			wantErr: false,
			want:    testBudget.Sum + testProperty.Price,
		},
		{
			name: " Property does not exist",
			fields: fields{
				InvestorId: testInvestorId,
				PropertyId: "",
				Time:       time.Now(),
			},
			wantErr: true,
			want:    0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sell := &SellingProperty{
				InvestorId: tt.fields.InvestorId,
				PropertyId: tt.fields.PropertyId,
				Time:       tt.fields.Time,
			}
			err := sell.Sell()
			if (err != nil) != tt.wantErr {
				t.Errorf("SellingProperty.Sell() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			totalbudget := &Budget{
				InvestorId: tt.fields.InvestorId,
				Sum:        tt.want,
			}

			if !tt.wantErr {
				updatedBudget, _ := totalbudget.Get()
				if updatedBudget.Sum != tt.want {
					t.Errorf("SellingProperty.Sell() = %v, want %v", updatedBudget.Sum, tt.want)
				}
			}
		})
	}

}
