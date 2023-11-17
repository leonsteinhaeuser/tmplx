package template

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGenericFuncMap(t *testing.T) {
	got := GenericFuncMap()
	if _, ok := got["filterMap"]; !ok {
		t.Errorf("GenericFuncMap() = %v, want %v", got, "filterMap")
	}
}

func Test_filterMap(t *testing.T) {
	type args struct {
		data      map[string]any
		keyPrefix string
	}
	tests := []struct {
		name string
		args args
		want map[string]any
	}{
		{
			name: "empty",
			args: args{
				data:      map[string]any{},
				keyPrefix: "",
			},
			want: map[string]any{},
		},
		{
			name: "empty prefix",
			args: args{
				data: map[string]any{
					"key1": "value1",
					"key2": "value2",
				},
				keyPrefix: "",
			},
			want: map[string]any{
				"key1": "value1",
				"key2": "value2",
			},
		},
		{
			name: "non-empty prefix",
			args: args{
				data: map[string]any{
					"key1": "value1",
					"key2": "value2",
				},
				keyPrefix: "key",
			},
			want: map[string]any{
				"key1": "value1",
				"key2": "value2",
			},
		},
		{
			name: "non-empty prefix with one match",
			args: args{
				data: map[string]any{
					"key1": "value1",
					"key2": "value2",
				},
				keyPrefix: "key1",
			},
			want: map[string]any{
				"key1": "value1",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			diff := cmp.Diff(filterMap(tt.args.data, tt.args.keyPrefix), tt.want)
			if diff != "" {
				t.Errorf("filterMap() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func Test_mapLength(t *testing.T) {
	type args struct {
		dt map[string]any
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "empty",
			args: args{
				dt: map[string]any{},
			},
			want: 0,
		},
		{
			name: "non-empty",
			args: args{
				dt: map[string]any{
					"key1": "value1",
					"key2": "value2",
				},
			},
			want: 2,
		},
		{
			name: "nil",
			args: args{
				dt: nil,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mapLength(tt.args.dt); got != tt.want {
				t.Errorf("mapLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
