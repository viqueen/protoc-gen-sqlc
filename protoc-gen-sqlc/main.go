package main

import (
	"github.com/viqueen/protoc-gen-sqlc/protoc-gen-sqlc/handler"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/pluginpb"
	"io"
	"log"
	"os"
)

func main() {
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("failed to read from stdin: %v", err)
	}
	request := &pluginpb.CodeGeneratorRequest{}
	if err = proto.Unmarshal(data, request); err != nil {
		log.Fatalf("failed to unmarshal input: %v", err)
	}

	response := &pluginpb.CodeGeneratorResponse{}
	for _, protoFile := range request.GetProtoFile() {
		err = handler.ProtoFileHandler(protoFile, response)
		if err != nil {
			response.Error = proto.String(err.Error())
		}
	}
	respond(response)
}

func respond(resp *pluginpb.CodeGeneratorResponse) {
	out, err := proto.Marshal(resp)
	if err != nil {
		log.Fatalf("Failed to marshal response: %v", err)
	}
	_, err = os.Stdout.Write(out)
	if err != nil {
		log.Fatalf("Failed to write response: %v", err)
	}
}
