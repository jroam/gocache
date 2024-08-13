package gocache

import (
	"testing"
)

func Test_Get(t *testing.T) {
	type args struct {
		key interface{}
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			name: "Get existing key",
			args: args{key: "key"},
			want: "value1",
		},
		{
			name: "Get non existing key",
			args: args{key: "non_existing_key"},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Delete(tt.args.key)
			Set(tt.args.key, tt.want)

			if got := Get(tt.args.key); got != tt.want {
				t.Errorf("key:=>%s,  %v != want %v", tt.args.key, got, tt.want)
			}
		})
	}
}
