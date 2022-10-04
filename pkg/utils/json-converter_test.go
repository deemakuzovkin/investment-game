package utils

import (
	"reflect"
	"testing"
)

func TestToJsonBytes(t *testing.T) {
	type args struct {
		input interface{}
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "int64",
			args: args{
				input: int64(1000000000000),
			},
			want: []byte{49, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48},
		},
		{
			name: "int64",
			args: args{
				input: map[string]interface{}{
					"name": "Golang",
				},
			},
			want: []byte{123, 34, 110, 97, 109, 101, 34, 58, 34, 71, 111, 108, 97, 110, 103, 34, 125},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToJsonBytes(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToJsonBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}
