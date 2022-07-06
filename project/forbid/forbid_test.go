package forbid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewForbidden(t *testing.T) {
	forbid := NewForbidden()
	assert.NotNil(t, forbid)
}

func TestForbidden_Insert(t *testing.T) {
	forbid := NewForbidden()

	strSlice := []string{
		" 熟透",
		"英文版 ",
		"女",
		"壮 夫",
		"壮汉",
	}

	for _, str := range strSlice {
		forbid.Insert(str)
	}
}

func TestForbidden_Search(t *testing.T) {
	forbid := NewForbidden()
	strSlice := []string{
		" 熟透",
		"英文版 ",
		"女",
		"女人",
		"壮 夫",
		"壮汉",
	}

	type args struct {
		word string
	}

	exactMatchSearch := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "", args: args{word: " 熟透 "}, want: true,
		},
		{
			name: "", args: args{word: " 英文版 "}, want: true,
		},
		{
			name: "", args: args{word: " 女 "}, want: true,
		},
		{
			name: "", args: args{word: " 女人1 "}, want: false,
		},
		{
			name: "", args: args{word: " 1壮 夫 "}, want: false,
		},
		{
			name: "", args: args{word: " 壮 夫 "}, want: true,
		},
		{
			name: "", args: args{word: " 壮"}, want: false,
		},
	}

	for _, str := range strSlice {
		forbid.Insert(str)
	}

	for _, tt := range exactMatchSearch {
		t.Run(tt.name, func(t *testing.T) {
			str, ok := forbid.ExactMatchSearch(tt.args.word)
			if ok != tt.want {
				t.Errorf("ExactMatchSearch() = %v, want %v", ok, tt.want)
			}
			t.Log(str)
		})
	}
}

func TestForbidden_CommonPrefixSearch(t *testing.T) {
	forbid := NewForbidden()
	strSlice := []string{
		" 熟透",
		"英文版 ",
		"女",
		"女人",
		"壮 夫",
		"壮汉",
		"你好菜",
	}

	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "", args: args{str: " 1熟透 "}, want: true,
		},
		{
			name: "", args: args{str: " 熟透11 "}, want: true,
		},
		{
			name: "", args: args{str: " 111熟透 "}, want: true,
		},
		{
			name: "", args: args{str: "  女"}, want: true,
		},
		{
			name: "", args: args{str: " 女|人 "}, want: true,
		},
		{
			name: "", args: args{str: " 壮 夫 "}, want: true,
		},
		{
			name: "", args: args{str: " 壮 夫| "}, want: true,
		},
		{
			name: "", args: args{str: " 111熟透 女人壮壮 夫"}, want: true,
		},
		{
			name: "", args: args{str: " 你好"}, want: false,
		},
	}

	for _, str := range strSlice {
		forbid.Insert(str)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			str, ok := forbid.CommonPrefixSearch(tt.args.str)
			if ok != tt.want {
				t.Errorf("CommonPrefixSearch() got = %v, want %v", ok, tt.want)
			}
			t.Log(str)
		})
	}
}
