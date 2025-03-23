package template

import (
	"bytes"
	"text/template"
)

func RenderTemplate(content string, data map[string]interface{}, strict bool) (string, error) {
	opt := "missingkey=default"
	if strict {
		opt = "missingkey=error"
	}
	tmpl, err := template.New("tpl").Option(opt).Parse(content)
	if err != nil {
		return "", err
	}
	var out bytes.Buffer
	err = tmpl.Execute(&out, data)
	return out.String(), err
}
