package codegen

import (
	"bytes"
	"google.golang.org/protobuf/types/descriptorpb"
	"strings"
	"text/template"
)

func SQLSchemaFile(message *descriptorpb.DescriptorProto, tableName string) (string, error) {
	tmpl, err := template.New("sqlSchemaFile").Funcs(template.FuncMap{
		"add": func(a, b int) int { return a + b },
	}).Parse(sqlSchemaFileTemplate)
	if err != nil {
		return "", err
	}
	params := extractSQLSchemaFileParams(message, tableName)
	var buffer bytes.Buffer
	err = tmpl.Execute(&buffer, params)
	if err != nil {
		return "", err
	}
	return buffer.String(), nil
}

type sqlSchemaFileParams struct {
	TableName    string
	Columns      []sqlColumn
	ForeignKeys  []sqlForeignKey
	ColumnsCount int
}

type sqlColumn struct {
	Name       string
	SQLType    string
	Constraint string
}

type sqlForeignKey struct {
	ColumnName       string
	ReferencesTable  string
	ReferencesColumn string
}

func extractSQLSchemaFileParams(message *descriptorpb.DescriptorProto, tableName string) sqlSchemaFileParams {
	var columns []sqlColumn
	var foreignKeys []sqlForeignKey
	for _, field := range message.GetField() {
		column := sqlColumn{
			Name: field.GetName(),
		}

		// TODO: handle more field types
		switch field.GetType() {
		case descriptorpb.FieldDescriptorProto_TYPE_STRING:
			column.SQLType = "VARCHAR(255)"
		default:
			column.SQLType = "TEXT"
		}

		// TODO: handle field options instead
		// Determine the constraints
		if field.GetName() == "id" {
			column.Constraint = "PRIMARY KEY"
			column.SQLType = "UUID"
		}

		// TODO: handle field options instead
		// Determine the foreign keys
		if strings.HasSuffix(field.GetName(), "_id") {
			column.SQLType = "UUID"
			column.Constraint = "NOT NULL"
			referencesTable := strings.TrimSuffix(field.GetName(), "_id")
			foreignKeys = append(foreignKeys, sqlForeignKey{
				ColumnName:       field.GetName(),
				ReferencesTable:  referencesTable,
				ReferencesColumn: "id",
			})
		}

		columns = append(columns, column)
	}
	return sqlSchemaFileParams{
		TableName:    tableName,
		Columns:      columns,
		ForeignKeys:  foreignKeys,
		ColumnsCount: len(columns),
	}
}

var sqlSchemaFileTemplate = `
-- This file was generated by protoc-gen-sqlc. DO NOT EDIT.

CREATE TABLE IF NOT EXISTS {{ .TableName }}
(
{{- range $index, $column := .Columns }}
    {{$column.Name}} {{$column.SQLType}} {{ $column.Constraint }}{{ if ne (add $index 1) $.ColumnsCount }},{{ end }}
{{- end }}
{{- range .ForeignKeys }}
    , FOREIGN KEY ({{ .ColumnName }}) REFERENCES {{ .ReferencesTable }} ({{ .ReferencesColumn }})
{{- end }}

	, created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMP WITH TIME ZONE
);
`
