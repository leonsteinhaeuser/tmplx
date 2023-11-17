package source

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestJSONFile(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]any
		wantErr bool
	}{
		{
			name: "want 4",
			args: args{
				path: TestJSONFilePath,
			},
			want: map[string]any{
				TestEnvKeyPrefix + "KEY1":    "value1",
				TestEnvKeyPrefix + "KEY2":    "value2",
				TestEnvKeySubPrefix + "KEY1": "value1",
				TestEnvKeySubPrefix + "KEY2": "value2",
			},
			wantErr: false,
		},
		{
			name: "unable to parse",
			args: args{
				path: "json.go",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "file not found",
			args: args{
				path: "not-found.json",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := JSONFile(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("JSONFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			diff := cmp.Diff(got, tt.want)
			if diff != "" {
				t.Errorf("JSONFile() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
