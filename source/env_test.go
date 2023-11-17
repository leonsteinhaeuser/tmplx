package source

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestEnviron(t *testing.T) {
	type args struct {
		prefix string
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
				prefix: "GO_TEST_ENV_",
			},
			want: map[string]any{
				TestEnvKeyPrefix + "KEY1":    "value1",
				TestEnvKeyPrefix + "KEY2":    "value2",
				TestEnvKeySubPrefix + "KEY1": "value1",
				TestEnvKeySubPrefix + "KEY2": "value2",
			},
		},
		{
			name: "want 2",
			args: args{
				prefix: "GO_TEST_ENV_PREFIX_SUB_",
			},
			want: map[string]any{
				TestEnvKeySubPrefix + "KEY1": "value1",
				TestEnvKeySubPrefix + "KEY2": "value2",
			},
		},
		{
			name: "want 0",
			args: args{
				prefix: "GO_TEST_ENV_PREFIX_SUB_SUB_",
			},
			want: map[string]any{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Environ(tt.args.prefix)
			if (err != nil) != tt.wantErr {
				t.Errorf("Environ() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			diff := cmp.Diff(got, tt.want)
			if diff != "" {
				t.Errorf("Environ() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
