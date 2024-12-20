package generator

import (
	"archshell/internal/pkg/batch"
	"bytes"
	"fmt"
	"html/template"
	"path"

	"gopkg.in/yaml.v2"
)

func Load(file string, context interface{}, show bool) (*batch.Batch, error) {

	name := path.Base(file)
	tmpl, err := template.New(name).ParseFiles(file)
	if err != nil {
		return nil, fmt.Errorf("construct template: %w", err)
	}

	var buf bytes.Buffer

	err = tmpl.Execute(&buf, context)
	if err != nil {
		return nil, fmt.Errorf("execute template: %w", err)
	}

	if show {
		fmt.Print(buf.String())
	}

	var u batch.Batch

	err = yaml.Unmarshal(buf.Bytes(), &u)
	if err != nil {
		return nil, fmt.Errorf("unmarshal batch: %w", err)
	}

	return &u, nil
}
