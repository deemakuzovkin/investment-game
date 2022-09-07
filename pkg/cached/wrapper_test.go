package cached

import (
	"git.mills.io/prologic/bitcask"
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
		db *bitcask.Bitcask
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
				db: testCacheData.db,
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
				db: testCacheData.db,
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
			data := &DataCache{
				db: tt.fields.db,
			}
			if err := data.Add(tt.args.key, tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("Add() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDataCache_Get(t *testing.T) {

	type fields struct {
		db *bitcask.Bitcask
	}
	type args struct {
		key []byte
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
				db: testCacheData.db,
			},
			args: args{
				key: testKey,
			},
			want:    testValue,
			wantErr: false,
		},
		{
			name: "get nil key",
			fields: fields{
				db: testCacheData.db,
			},
			args: args{
				key: nil,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := &DataCache{
				db: tt.fields.db,
			}
			got, err := data.Get(tt.args.key)
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
