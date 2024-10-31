package handler_test

import (
	"context"
	"github.com/bufbuild/protocompile"
	"github.com/bufbuild/protocompile/protoutil"
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
	compiler := protocompile.Compiler{
		Resolver: &protocompile.SourceResolver{
			ImportPaths: []string{
				filepath.Join(cwd, "../../test-protos"),
				filepath.Join(cwd, "../../protos"),
			},
		},
	}
	descriptors, err := compiler.Compile(context.Background(), musicProto)
	require.NoError(t, err)

	response := &pluginpb.CodeGeneratorResponse{}
	for _, desc := range descriptors {
		protoDescriptor := protoutil.ProtoFromFileDescriptor(desc)
		err = handler.ProtoFileHandler(protoDescriptor, response)
		assert.NoError(t, err)
		assert.Len(t, response.File, 2)
	}
}
