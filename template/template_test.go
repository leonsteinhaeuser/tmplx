package template

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestDataSource(t *testing.T) {
	type args struct {
		src string
	}
	tests := []struct {
		name     string
		args     args
		wantFunc bool
		wantErr  bool
	}{
		{
			name: "env",
			args: args{
				src: "env",
			},
			wantFunc: true,
			wantErr:  false,
		},
		{
			name: "json",
			args: args{
				src: "json",
			},
			wantFunc: true,
			wantErr:  false,
		},
		{
			name: "yaml",
			args: args{
				src: "yaml",
			},
			wantFunc: true,
			wantErr:  false,
		},
		{
			name: "unsupported",
			args: args{
				src: "unsupported",
			},
			wantFunc: false,
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DataSource(tt.args.src)
			if (err != nil) != tt.wantErr {
				t.Errorf("DataSource() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantFunc && got == nil {
				t.Errorf("DataSource() = %v, want a function", got)
				return
			}
		})
	}
}

func Test_newTemplate(t *testing.T) {
	type args struct {
		templateData []byte
	}
	tests := []struct {
		name            string
		args            args
		wantNilTemplate bool
		wantErr         bool
	}{
		{
			name: "test",
			args: args{
				templateData: []byte(`Hello {{ .Values.name }}!`),
			},
			wantNilTemplate: false,
			wantErr:         false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := newTemplate(tt.args.templateData)
			if (err != nil) != tt.wantErr {
				t.Errorf("newTemplate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantNilTemplate && got != nil {
				t.Errorf("newTemplate() = %v, want nil", got)
				return
			}
		})
	}
}

func TestRender(t *testing.T) {
	type args struct {
		templateData []byte
		data         map[string]any
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "simple",
			args: args{
				templateData: []byte(`Hello {{ .Values.name }}!`),
				data: map[string]any{
					"name": "World",
				},
			},
			want:    []byte(`Hello World!`),
			wantErr: false,
		},
		{
			name: "complex",
			args: args{
				templateData: []byte(`Hello {{ .Values.name }}! {{ .Values.name }} is {{ .Values.age }} years old.`),
				data: map[string]any{
					"name": "World",
					"age":  42,
				},
			},
			want:    []byte(`Hello World! World is 42 years old.`),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Render(tt.args.templateData, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("Render() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			diff := cmp.Diff(got, tt.want)
			if diff != "" {
				t.Errorf("Render() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
