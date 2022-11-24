package models

import (
	"github.com/google/uuid"
	"testing"
)

func TestInvestor_AddOrUpdate(t *testing.T) {
	type fields struct {
		Id      string
		Name    string
		Surname string
		Age     int
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "add investor",
			fields: fields{
				Id:      uuid.New().String(),
				Name:    "Test investor",
				Surname: "Test surname",
				Age:     34,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inv := &Investor{
				Id:      tt.fields.Id,
				Name:    tt.fields.Name,
				Surname: tt.fields.Surname,
				Age:     tt.fields.Age,
			}
			if err := inv.AddOrUpdate(); (err != nil) != tt.wantErr {
				t.Errorf("AddOrUpdate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestInvestor_Get(t *testing.T) {
	type fields struct {
		Id      string
		Name    string
		Surname string
		Age     int
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "get investor by key",
			fields: fields{
				Id:      uuid.New().String(),
				Name:    "Test investor by Id",
				Surname: "Some surname",
				Age:     23,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inv := &Investor{
				Id:      tt.fields.Id,
				Name:    tt.fields.Name,
				Surname: tt.fields.Surname,
				Age:     tt.fields.Age,
			}
			err := inv.AddOrUpdate()
			if err != nil {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			gotInvestor, err := inv.Get()
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotInvestor.Id != inv.Id {
				t.Errorf("Not equals inputId = %v, outId %v", inv.Id, gotInvestor.Id)
			}
		})
	}
}
