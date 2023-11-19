package template

import (
	"testing"
	"text/template"
)

func TestWithDelims(t *testing.T) {
	type args struct {
		left  string
		right string
	}
	tests := []struct {
		name string
		args args
		want TemplateOption
	}{
		{
			name: "basic check if delimiters can be set",
			args: args{
				left:  "{{",
				right: "}}",
			},
			want: func(t *template.Template) {
				t.Delims("{{", "}}")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t1 := template.New("test")
			WithDelims(tt.args.left, tt.args.right)(t1)
			// TODO: check if delimeters are set
		})
	}
}
