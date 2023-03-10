package app

import (
	"reflect"
	"testing"
)

func Test_getApps(t *testing.T) {
	tests := []struct {
		name string
		want Apps
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getApps(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getApps() = %v, want %v", got, tt.want)
			}
		})
	}
}
