package helpers

import (
	"fmt"
	sqlcv1 "github.com/viqueen/protoc-gen-sqlc/api/sqlc/v1"
	"google.golang.org/protobuf/types/descriptorpb"
	"strings"
)

func SqlcEntityOption(message *descriptorpb.DescriptorProto) (string, bool) {
	return hasOption(message, fmt.Sprintf("[%s]", sqlcv1.E_SqlcEntity.TypeDescriptor().FullName()))
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
