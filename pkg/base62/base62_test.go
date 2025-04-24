package base62

import (
	"testing"
)

func TestUint2string(t *testing.T) {
	type args struct {
		seq uint64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1", args: args{seq: 1}, want: "1"},
		{name: "10", args: args{seq: 10}, want: "a"},
		{name: "62", args: args{seq: 62}, want: "10"},
		{name: "65530", args: args{seq: 65530}, want: "h2W"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Uint2string(tt.args.seq); got != tt.want {
				t.Errorf("Uint2string() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestString2Uint(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{name: "a", args: args{"a"}, want: 10},
		{name: "11", args: args{"11"}, want: 63},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := String2Uint(tt.args.s); got != tt.want {
				t.Errorf("String2Uint() = %v, want %v", got, tt.want)
			}
		})
	}
}
