package gocache

import (
	"fmt"
	"testing"
)

type TestStruct struct {
	Name string `json:"name"`
}

func Test_Type(t *testing.T) {
	a1 := &TestStruct{Name: "test!!"}
	fmt.Println(a1.Name)
}

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
		{
			name: "Get non existing key",
			args: args{key: "list"},
			want: []int{1, 2, 3},
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
