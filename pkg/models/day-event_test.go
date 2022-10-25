package models

import (
	"github.com/google/uuid"
	"reflect"
	"testing"
	"time"
)

var (
	testEvent = &DayEvent{
		Id:      "111-222-333-444-555",
		Time:    time.Date(2022, time.October, 25, 20, 17, 0, 0, time.UTC),
		Name:    "Test name",
		Comment: "Test comment",
		Kf:      1.2,
	}
)

func TestDayEvent_AddOrUpdate(t *testing.T) {
	type fields struct {
		Id      string
		Time    time.Time
		Name    string
		Comment string
		Kf      float64
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "add day event",
			fields: fields{
				Id:      uuid.New().String(),
				Time:    time.Now(),
				Name:    "Доходность...",
				Comment: "Test event",
				Kf:      1.1,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			prop := &DayEvent{
				Id:      tt.fields.Id,
				Time:    tt.fields.Time,
				Name:    tt.fields.Name,
				Comment: tt.fields.Comment,
				Kf:      tt.fields.Kf,
			}
			if err := prop.AddOrUpdate(); (err != nil) != tt.wantErr {
				t.Errorf("AddOrUpdate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDayEvent_Get(t *testing.T) {
	type fields struct {
		Id      string
		Time    time.Time
		Name    string
		Comment string
		Kf      float64
	}
	tests := []struct {
		name    string
		fields  fields
		want    *DayEvent
		wantErr bool
	}{
		{
			name: "get by key",
			fields: fields{
				Id:      testEvent.Id,
				Time:    testEvent.Time,
				Name:    testEvent.Name,
				Comment: testEvent.Comment,
				Kf:      testEvent.Kf,
			},
			want:    testEvent,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			prop := &DayEvent{
				Id:      tt.fields.Id,
				Time:    tt.fields.Time,
				Name:    tt.fields.Name,
				Comment: tt.fields.Comment,
				Kf:      tt.fields.Kf,
			}
			if err := prop.AddOrUpdate(); (err != nil) != tt.wantErr {
				t.Errorf("AddOrUpdate() error = %v, wantErr %v", err, tt.wantErr)
			}
			got, err := prop.Get()
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

func TestDayEvent_GetLatest(t *testing.T) {
	type fields struct {
		Id      string
		Time    time.Time
		Name    string
		Comment string
		Kf      float64
	}
	tests := []struct {
		name    string
		fields  fields
		want    *DayEvent
		wantErr bool
	}{
		{
			name: "get latest",
			fields: fields{
				Id:      testEvent.Id,
				Time:    testEvent.Time,
				Name:    testEvent.Name,
				Comment: testEvent.Comment,
				Kf:      testEvent.Kf,
			},
			want:    testEvent,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			prop := &DayEvent{
				Id:      tt.fields.Id,
				Time:    tt.fields.Time,
				Name:    tt.fields.Name,
				Comment: tt.fields.Comment,
				Kf:      tt.fields.Kf,
			}
			if err := prop.AddOrUpdate(); (err != nil) != tt.wantErr {
				t.Errorf("AddOrUpdate() error = %v, wantErr %v", err, tt.wantErr)
			}
			got, err := prop.GetLatest()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetLatest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetLatest() got = %v, want %v", got, tt.want)
			}
		})
	}
}
