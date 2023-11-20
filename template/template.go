package template

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/Masterminds/sprig/v3"
	"github.com/leonsteinhaeuser/tmplx/source"
)

const (
	formatKeyEnv  = "env"
	formatKeyJSON = "json"
	formatKeyYAML = "yaml"
)

type DataSourceFunc func(string) (map[string]any, error)

var (
	ErrUnsupportedFormat = fmt.Errorf("unsupported format")
)

// dataFunc returns the source function for the given format.
func DataSource(src string) (DataSourceFunc, error) {
	switch src {
	case formatKeyEnv:
		return source.Environ, nil
	case formatKeyJSON:
		return source.JSONFile, nil
	case formatKeyYAML:
		return source.YAMLFile, nil
	default:
		return nil, fmt.Errorf("%w: %s", ErrUnsupportedFormat, src)
	}
}

func newTemplate(templateData []byte, options ...TemplateOption) (*template.Template, error) {
	t := template.New("my").Funcs(sprig.FuncMap()).Funcs(GenericFuncMap())
	// apply options
	for _, opt := range options {
		opt(t)
	}
	return t.Parse(string(templateData))
}

// Render renders the template data to the given writer.
func Render(templateData []byte, data map[string]any, options ...TemplateOption) ([]byte, error) {
	// create template
	t, err := newTemplate(templateData, options...)
	if err != nil {
		return nil, err
	}
	// execute template
	bbuf := bytes.NewBuffer([]byte{})
	err = t.Execute(bbuf, Data{Values: data})
	if err != nil {
		return nil, err
	}
	return bbuf.Bytes(), nil
}
