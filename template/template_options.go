package template

import (
	"text/template"
)

type TemplateOption func(*template.Template)

// WithDelims returns a template option that sets the left and right delimiters.
func WithDelims(left, right string) TemplateOption {
	return func(t *template.Template) {
		// overwrite delimiters
		t.Delims(left, right)
	}
}
