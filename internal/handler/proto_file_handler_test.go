package handler_test

import (
	"github.com/jhump/protoreflect/desc/protoparse"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/viqueen/protoc-gen-sqlc/internal/handler"
	"google.golang.org/protobuf/types/pluginpb"
	"os"
	"path/filepath"
	"testing"
)

func TestProtoFileHandler(t *testing.T) {
	t.Parallel()

	cwd, err := os.Getwd()
	require.NoError(t, err)

	musicProto := "music/v1/music_models.proto"
	parser := protoparse.Parser{
		ImportPaths: []string{
			filepath.Join(cwd, "../../test-protos"),
			filepath.Join(cwd, "../../protos"),
		},
	}
	descriptors, err := parser.ParseFiles(musicProto)
	require.NoError(t, err)

	response := &pluginpb.CodeGeneratorResponse{}
	for _, desc := range descriptors {
		err = handler.ProtoFileHandler(desc.AsFileDescriptorProto(), response)
		assert.NoError(t, err)
		assert.Len(t, response.File, 2)
	}
}
