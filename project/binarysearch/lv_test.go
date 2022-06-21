package binarysearch

import "testing"

func TestUseSub(t *testing.T) {
	type args struct {
		exp int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{exp: 0},
			want: 1,
		},
		{
			args: args{exp: 50},
			want: 1,
		},
		{
			args: args{exp: 601},
			want: 13,
		},
		{
			args: args{exp: 10500},
			want: 20,
		},
		{
			args: args{exp: 20500},
			want: 20,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UseSub(tt.args.exp); got != tt.want {
				t.Errorf("UseSub() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUseBinarySearch(t *testing.T) {
	type args struct {
		exp int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{exp: 0},
			want: 1,
		},
		{
			args: args{exp: 50},
			want: 1,
		},
		{
			args: args{exp: 601},
			want: 13,
		},
		{
			args: args{exp: 10500},
			want: 20,
		},
		{
			args: args{exp: 20500},
			want: 20,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UseBinarySearch(tt.args.exp); got != tt.want {
				t.Errorf("UseBinarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
