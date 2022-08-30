package cached

import (
	"testing"
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
