package handler

import (
	"fmt"
	sqlcv1 "github.com/viqueen/protoc-gen-sqlc/api/sqlc/v1"
	"github.com/viqueen/protoc-gen-sqlc/pkg/codegen"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/pluginpb"
	"strings"
)

func ProtoFileHandler(protoFile *descriptorpb.FileDescriptorProto, response *pluginpb.CodeGeneratorResponse) error {
	messages := protoFile.GetMessageType()
	if len(messages) == 0 {
		return nil
	}
	migrationIndex := 0
	for _, message := range messages {
		tableName, ok := sqlcEntityOption(message)
		if !ok {
			continue
		}
		trimmedTableName := strings.Trim(tableName, "\"")
		sqlSchemaFileName := fmt.Sprintf("V%04d__%s_table.sql", migrationIndex, trimmedTableName)
		sqlSchemaFileContent, err := codegen.SQLSchemaFile(message, trimmedTableName)
		if err != nil {
			return err
		}
		response.File = append(response.File, &pluginpb.CodeGeneratorResponse_File{
			Name:    &sqlSchemaFileName,
			Content: &sqlSchemaFileContent,
		})
	}
	return nil
}

func sqlcEntityOption(message *descriptorpb.DescriptorProto) (string, bool) {
	return hasOption(message, fmt.Sprintf("[%s]", sqlcv1.E_SqlcEntity.Name))
}

func hasOption(message *descriptorpb.DescriptorProto, option string) (string, bool) {
	options := message.GetOptions()
	if options == nil {
		return "", false
	}
	optionsMap := parseOptions(options.String())
	value, ok := optionsMap[option]
	return value, ok
}

func parseOptions(options string) map[string]string {
	optionsMap := make(map[string]string)
	tokens := strings.Split(options, " ")
	for _, token := range tokens {
		parts := strings.Split(token, ":")
		optionsMap[parts[0]] = parts[1]
	}
	return optionsMap
}
