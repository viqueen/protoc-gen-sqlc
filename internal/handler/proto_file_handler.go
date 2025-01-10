package handler

import (
	"fmt"
	"github.com/viqueen/protoc-gen-sqlc/internal/codegen"
	"github.com/viqueen/protoc-gen-sqlc/pkg/helpers"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/pluginpb"
	"path/filepath"
	"strings"
)

func ProtoFileHandler(protoFile *descriptorpb.FileDescriptorProto, response *pluginpb.CodeGeneratorResponse) error {
	messages := protoFile.GetMessageType()
	if len(messages) == 0 {
		return nil
	}
	migrationIndex := 0
	for _, message := range messages {
		tableName, ok := helpers.SqlcEntityOption(message)
		if !ok {
			continue
		}
		trimmedTableName := strings.Trim(tableName, "\"")
		sqlSchemaFileName := fmt.Sprintf("V%04d__%s_table.sql", migrationIndex, trimmedTableName)
		sqlSchemaFilePath := filepath.Join("data", "schema", sqlSchemaFileName)
		sqlSchemaFileContent, err := codegen.SQLSchemaFile(message, trimmedTableName)
		if err != nil {
			return err
		}
		response.File = append(response.File, &pluginpb.CodeGeneratorResponse_File{
			Name:    proto.String(sqlSchemaFilePath),
			Content: proto.String(sqlSchemaFileContent),
		})
		migrationIndex++

		sqlQueriesFileName := fmt.Sprintf("%s_queries.sql", trimmedTableName)
		sqlQueriesFilePath := filepath.Join("data", "queries", sqlQueriesFileName)
		sqlQueriesFileContent, err := codegen.SQLQueriesFile(message, trimmedTableName)
		if err != nil {
			return err
		}
		response.File = append(response.File, &pluginpb.CodeGeneratorResponse_File{
			Name:    proto.String(sqlQueriesFilePath),
			Content: proto.String(sqlQueriesFileContent),
		})
	}
	return nil
}
