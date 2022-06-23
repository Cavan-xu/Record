package factory

import (
	"reflect"
	"testing"
)

func TestNewConfigParseFactory(t *testing.T) {
	type args struct {
		c string
	}
	tests := []struct {
		name            string
		args            args
		wantFactory     ConfigParseFactory
		wantConfigParse ConfigParse
	}{
		{
			name:            "test json factory",
			args:            args{c: "json"},
			wantFactory:     JsonConfigParseFactory{},
			wantConfigParse: JsonConfigParse{},
		},
		{
			name:            "test xml factory",
			args:            args{c: "xml"},
			wantFactory:     XmlConfigParseFactory{},
			wantConfigParse: XmlConfigParser{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotFactory := NewConfigParseFactory(tt.args.c)
			if !reflect.DeepEqual(gotFactory, tt.wantFactory) {
				t.Errorf("NewConfigParseFactory() = %v, want %v", gotFactory, tt.wantFactory)
			}
			if gotConfigParse := gotFactory.NewConfigParse(); !reflect.DeepEqual(gotConfigParse, tt.wantConfigParse) {
				t.Errorf("NewConfigParse() = %v, want %v", gotConfigParse, tt.wantConfigParse)
			}
		})
	}
}
