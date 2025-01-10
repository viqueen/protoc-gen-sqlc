package helpers

import (
	"fmt"
	"google.golang.org/protobuf/types/descriptorpb"
	sqlcv1 "protoc-gen-sqlc/api/sqlc/v1"
	"strings"
)

func SqlcEntityOption(message *descriptorpb.DescriptorProto) (string, bool) {
	return hasOption(message, fmt.Sprintf("[%s]", sqlcv1.E_SqlcEntity.TypeDescriptor().FullName()))
}

func SqlcRequestOption(message *descriptorpb.DescriptorProto) (string, bool) {
	return hasOption(message, fmt.Sprintf("[%s]", sqlcv1.E_SqlcRequest.TypeDescriptor().FullName()))
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
