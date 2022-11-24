package models

import (
	"github.com/google/uuid"
	"reflect"
	"testing"
)

func TestBudget_AddOrUpdate(t *testing.T) {
	type fields struct {
		InvestorId string
		Sum        float64
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "add budget",
			fields: fields{
				InvestorId: uuid.New().String(),
				Sum:        1000.0,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			budget := &Budget{
				InvestorId: tt.fields.InvestorId,
				Sum:        tt.fields.Sum,
			}
			if err := budget.AddOrUpdate(); (err != nil) != tt.wantErr {
				t.Errorf("AddOrUpdate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBudget_Get(t *testing.T) {
	testBudget := Budget{
		InvestorId: uuid.New().String(),
		Sum:        1200.0,
	}
	type fields struct {
		InvestorId string
		Sum        float64
	}
	tests := []struct {
		name    string
		fields  fields
		want    *Budget
		wantErr bool
	}{
		{
			name: "get budget by id",
			fields: fields{
				InvestorId: testBudget.InvestorId,
				Sum:        testBudget.Sum,
			},
			want:    &testBudget,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			budget := &Budget{
				InvestorId: tt.fields.InvestorId,
				Sum:        tt.fields.Sum,
			}
			if err := budget.AddOrUpdate(); (err != nil) != tt.wantErr {
				t.Errorf("AddOrUpdate() error = %v, wantErr %v", err, tt.wantErr)
			}
			got, err := budget.Get()
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
		})
	}
}
