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
		assert.Len(t, response.File, 4)
		assert.Equal(t, "data/schema/V0000__album_table.sql", response.File[0].GetName())
		assert.Equal(t, "data/queries/album_queries.sql", response.File[1].GetName())
		assert.Equal(t, "data/schema/V0001__track_table.sql", response.File[2].GetName())
		assert.Equal(t, "data/queries/track_queries.sql", response.File[3].GetName())
	}
}
