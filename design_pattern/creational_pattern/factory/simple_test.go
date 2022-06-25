package factory

import (
	"reflect"
	"testing"
)

func TestNewConfigParse(t *testing.T) {
	type args struct {
		c string
	}
	tests := []struct {
		name string
		args args
		want ConfigParse
	}{
		{
			name: "test new json",
			args: args{c: "json"},
			want: JsonConfigParse{},
		},
		{
			name: "test new xml",
			args: args{c: "xml"},
			want: XmlConfigParser{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewConfigParse(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewConfigParse() = %v, want %v", got, tt.want)
			}
		})
	}
}
