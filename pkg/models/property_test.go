package models

import (
	"github.com/google/uuid"
	"testing"
)

func TestProperty_AddOrUpdate(t *testing.T) {
	type fields struct {
		Id   string
		Name string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "add property",
			fields: fields{
				Id:   uuid.New().String(),
				Name: "Test property",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			prop := &Property{
				Id:   tt.fields.Id,
				Name: tt.fields.Name,
			}
			if err := prop.AddOrUpdate(); (err != nil) != tt.wantErr {
				t.Errorf("AddOrUpdate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestProperty_Get(t *testing.T) {
	type fields struct {
		Id   string
		Name string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "get property by name",
			fields: fields{
				Id:   uuid.New().String(),
				Name: "Test property by Id",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			prop := &Property{
				Id:   tt.fields.Id,
				Name: tt.fields.Name,
			}
			err := prop.AddOrUpdate()
			if err != nil {
				t.Errorf("AddOrUpdate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			gotProperty, err := prop.Get()
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotProperty.Id != prop.Id {
				t.Errorf("Not equals inputId = %v, outId %v", prop.Id, gotProperty.Id)
			}
		})
	}
}
