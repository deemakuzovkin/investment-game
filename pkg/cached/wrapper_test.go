package cached

import (
	"fmt"
	"github.com/google/uuid"
	"reflect"
	"testing"
)

const TestCachePath = "test-path"

var (
	testCacheData, _ = Connect(TestCachePath)
	testKey          = []byte("1")
	testValue        = []byte("1")
)

func TestConnect(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "connect",
			args: args{
				path: "test-path-cached",
			},
			wantErr: false,
		},
		{
			name: "connect",
			args: args{
				path: "second-path-cached",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Connect(tt.args.path)
			if (err != nil) != tt.wantErr && got == nil {
				t.Errorf("Connect() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestDataCache_Add(t *testing.T) {
	type fields struct {
		name string
	}
	type args struct {
		key   []byte
		value []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "add object to data",
			fields: fields{
				name: TestCachePath,
			},
			args: args{
				key:   testKey,
				value: testValue,
			},
			wantErr: false,
		},
		{
			name: "add nil object to data",
			fields: fields{
				name: TestCachePath,
			},
			args: args{
				key:   nil,
				value: []byte("2"),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dataCache, err := Connect(tt.fields.name)
			if (err != nil) != tt.wantErr && dataCache == nil {
				t.Errorf("Connect() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err := dataCache.Add(tt.args.key, tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("Add() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDataCache_Get(t *testing.T) {
	type fields struct {
		name string
	}
	type args struct {
		key   []byte
		value []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "get object",
			fields: fields{
				name: TestCachePath,
			},
			args: args{
				key:   testKey,
				value: testValue,
			},
			want:    testValue,
			wantErr: false,
		},
		{
			name: "get nil key",
			fields: fields{
				name: TestCachePath,
			},
			args: args{
				key:   nil,
				value: nil,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dataCache, err := Connect(tt.fields.name)
			if (err != nil) != tt.wantErr && dataCache == nil {
				t.Errorf("Connect() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err := dataCache.Add(tt.args.key, testValue); (err != nil) != tt.wantErr {
				t.Errorf("Add() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			got, err := dataCache.Get(tt.args.key)
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

func TestDataCache_Remove(t *testing.T) {
	type fields struct {
		name string
	}
	type args struct {
		key   []byte
		value []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "remove object",
			fields: fields{
				name: TestCachePath,
			},
			args: args{
				key:   testKey,
				value: testValue,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dataCache, err := Connect(tt.fields.name)
			if (err != nil) != tt.wantErr && dataCache == nil {
				t.Errorf("Connect() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err := dataCache.Add(tt.args.key, testValue); (err != nil) != tt.wantErr {
				t.Errorf("Add() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			getData, err := dataCache.Get(tt.args.key)
			if (err != nil) != tt.wantErr && getData == nil {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err := dataCache.Remove(tt.args.key); (err != nil) != tt.wantErr {
				t.Errorf("Remove() error = %v, wantErr %v", err, tt.wantErr)
			}
			getData, err = dataCache.Get(tt.args.key)
			if getData != nil || err.Error() != "error: key not found" {
				t.Errorf("Get after remove error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestDataCache_GetList(t *testing.T) {
	type fields struct {
		name  string
		count int
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "iterator",
			fields: fields{
				name:  uuid.New().String(),
				count: 100,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dataCache, err := Connect(tt.fields.name)
			if (err != nil) != tt.wantErr && dataCache == nil {
				t.Errorf("Connect() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			for i := 0; i < tt.fields.count; i++ {
				key := []byte(fmt.Sprintf("key-%d", i))
				value := []byte(fmt.Sprintf("value-%d", i))
				if err := dataCache.Add(key, value); (err != nil) != tt.wantErr {
					t.Errorf("Add() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
			}
			channelItems, err := dataCache.GetList()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			iterator := 0
			for item := range channelItems {
				if len(item.Key) <= 0 {
					t.Error("Key is empty")
					break
				}
				if len(item.Value) <= 0 {
					t.Error("Value is empty")
					break
				}
				iterator++
			}
			if tt.fields.count != iterator {
				t.Errorf("Count %v not equal %v", tt.fields.count, iterator)
				return
			}
		})
	}
}

func TestDataCache_Clear(t *testing.T) {
	type fields struct {
		name string
	}
	type args struct {
		key   []byte
		value []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "clear all data",
			args: args{
				key:   testKey,
				value: testValue,
			},
			fields: fields{
				name: uuid.New().String(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dataCache, err := Connect(tt.fields.name)
			if (err != nil) != tt.wantErr && dataCache == nil {
				t.Errorf("Connect() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err := dataCache.Add(tt.args.key, testValue); (err != nil) != tt.wantErr {
				t.Errorf("Add() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			getData, err := dataCache.Get(tt.args.key)
			if (err != nil) != tt.wantErr && getData == nil {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err := dataCache.Clear(); (err != nil) != tt.wantErr {
				t.Errorf("Clear() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
