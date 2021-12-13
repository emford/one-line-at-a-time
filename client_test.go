package main

import (
	"reflect"
	"testing"
)

func TestGetFile(t *testing.T) {
	tests := []struct {
		name    string
		want    []string
		wantErr bool
	}{
		{
			name: "success",
			want: []string{"Dunder Mifflin Paper Company, Inc., Scranton, PA",
				"Assistant [to the] Regional Manager",
				"Act as Regional Manager's eyes, ears, and right hand, overseeing...",
				"and reporting on employee conduct and productivity"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetFile()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
