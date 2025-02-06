package codegen

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"google.golang.org/protobuf/types/descriptorpb"
	"strings"
	"text/template"
)

func SQLQueriesFile(message *descriptorpb.DescriptorProto, tableName string) (string, error) {
	tmpl, err := template.New("sqlQueriesFile").Funcs(template.FuncMap{
		"join": func(slice []string, separator string) string {
			return strings.Join(slice, separator)
		},
		"title": cases.Title(language.English).String,
		"lower": strings.ToLower,
		"sub":   func(a, b int) int { return a - b },
	}).Parse(sqlQueriesFileTemplate)
	if err != nil {
		return "", err
	}
	params := extractSQLQueriesFileParams(message, tableName)
	var buf strings.Builder
	err = tmpl.Execute(&buf, params)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

type sqlQueriesFileParams struct {
	TableName   string
	MessageName string
	Columns     []string
	IdColumns   []string
}

func extractSQLQueriesFileParams(message *descriptorpb.DescriptorProto, tableName string) sqlQueriesFileParams {
	var columns []string
	var idColumns []string
	for _, field := range message.GetField() {
		columns = append(columns, field.GetName())
		if strings.HasSuffix(field.GetName(), "_id") {
			idColumns = append(idColumns, field.GetName())
		}
	}
	return sqlQueriesFileParams{
		TableName:   tableName,
		MessageName: message.GetName(),
		Columns:     columns,
		IdColumns:   idColumns,
	}
}

var sqlQueriesFileTemplate = `
-- name: Create{{ .MessageName }} :one
-- Create{{ .MessageName }} creates a new {{ .MessageName | title | lower }}.
INSERT INTO public.{{ .TableName }} ({{ join .Columns ", " }}) 
VALUES ({{ range $index, $column := .Columns }}@{{$column}}{{ if ne $index (sub (len $.Columns) 1)}} , {{ end }}{{ end }})
RETURNING *;

-- name: Update{{ .MessageName }} :one
-- Update{{ .MessageName }} updates a {{ .MessageName | title | lower }}.
UPDATE public.{{ .TableName }}
SET {{ range $index, $column := .Columns }}{{ $column }} = @{{$column}}{{ if ne $index (sub (len $.Columns) 1)}} , {{ end }}{{ end }}
WHERE id = @id
RETURNING *;

-- name: Delete{{ .MessageName }} :one
-- Delete{{ .MessageName }} deletes a {{ .MessageName | title | lower }}.
DELETE FROM public.{{ .TableName }}
WHERE id = @id
RETURNING *;

-- name: Get{{ .MessageName }} :one
-- Get{{ .MessageName }} gets a {{ .MessageName | title | lower }} by id.
SELECT * FROM public.{{ .TableName }}
WHERE id = @id;

-- name: List{{ .MessageName }}s :many
-- List{{ .MessageName }}s lists all {{ .MessageName | title | lower }}.
SELECT * FROM public.{{ .TableName }}
{{ if gt (len $.IdColumns) 0  }}
WHERE {{ range $index, $column := .IdColumns }}{{ $column }} = @{{ $column }}{{ if ne $index (sub (len $.IdColumns) 1)}} AND {{ end }}{{ end }}
{{ end }}
LIMIT @page_limit OFFSET @page_offset;
`
