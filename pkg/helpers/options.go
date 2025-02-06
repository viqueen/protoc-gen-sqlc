package helpers

import (
	"fmt"
	sqlcv1 "github.com/viqueen/protoc-gen-sqlc/api/sqlc/v1"
	"google.golang.org/protobuf/types/descriptorpb"
	"strings"
)

func SqlcEntityOption(message *descriptorpb.DescriptorProto) (string, bool) {
	return messageHasOption(message, fmt.Sprintf("[%s]", sqlcv1.E_SqlcEntity.TypeDescriptor().FullName()))
}

func SqlcRequestOption(message *descriptorpb.DescriptorProto) (string, bool) {
	return messageHasOption(message, fmt.Sprintf("[%s]", sqlcv1.E_SqlcRequest.TypeDescriptor().FullName()))
}

func SqlcFkOption(message *descriptorpb.FieldDescriptorProto) (string, bool) {
	return fieldHasOption(message, fmt.Sprintf("[%s]", sqlcv1.E_SqlcFk.TypeDescriptor().FullName()))
}

func fieldHasOption(field *descriptorpb.FieldDescriptorProto, option string) (string, bool) {
	options := field.GetOptions()
	if options == nil {
		return "", false
	}
	return hasOption(options.String(), option)
}

func messageHasOption(message *descriptorpb.DescriptorProto, option string) (string, bool) {
	options := message.GetOptions()
	if options == nil {
		return "", false
	}
	return hasOption(options.String(), option)
}

func hasOption(options string, option string) (string, bool) {
	optionsMap := parseOptions(options)
	value, ok := optionsMap[option]
	if !ok {
		return "", false
	}
	return strings.Trim(value, "\""), ok
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
