PROTO=proto

PROTO_GEN_GRPC 				:=
ifeq ($(OS),Windows_NT)
	PROTO_GEN_GRPC = $(CURDIR)/client/node_modules/.bin/grpc_tools_node_protoc_plugin.cmd
else
	PROTO_GEN_GRPC = $(CURDIR)/client/node_modules/.bin/grpc_tools_node_protoc_plugin
endif

all: proto-server proto-js

proto-server:
	protoc.exe --proto_path=proto --go_out=plugins=grpc:./server/proto ./proto/uploadfile.proto
	
proto-js:
	protoc.exe --proto_path=proto --js_out=import_style=commonjs,binary:./client/proto --grpc_out=./client/proto --plugin=protoc-gen-grpc=$(PROTO_GEN_GRPC) ./proto/uploadfile.proto
