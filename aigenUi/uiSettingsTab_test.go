package aigenUi

import (
	"reflect"
	"testing"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

func TestGenSettings(t *testing.T) {
	type args struct {
		mapungubwe fyne.App
	}
	tests := []struct {
		name string
		args args
		want *container.TabItem
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenSettings(tt.args.mapungubwe); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenSettings() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMultiModelSettings(t *testing.T) {
	type args struct {
		mapungubwe fyne.App
	}
	tests := []struct {
		name string
		args args
		want *container.TabItem
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MultiModelSettings(tt.args.mapungubwe); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MultiModelSettings() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAudioSettingsTab(t *testing.T) {
	tests := []struct {
		name string
		want *container.TabItem
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AudioSettingsTab(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AudioSettingsTab() = %v, want %v", got, tt.want)
			}
		})
	}
}
